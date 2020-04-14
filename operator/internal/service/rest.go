package service

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
	"github.com/fgituser/management-client-psychologist.services/operator/internal/transportclient"
	"github.com/fgituser/management-client-psychologist.services/operator/internal/transportpsychologist"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type restserver struct {
	router                chi.Router
	logger                *logrus.Logger
	userRoles             []*userRole
	transportPsychologist transportpsychologist.Transport
	transportClient       transportclient.Transport
}

func newRESTServer(router chi.Router, uRoles []*userRole,
	transportPsychologist transportpsychologist.Transport,
	transportClient transportclient.Transport) *restserver {
	r := &restserver{
		router:                router,
		logger:                logrus.New(),
		userRoles:             uRoles,
		transportPsychologist: transportPsychologist,
		transportClient:       transportClient,
	}

	r.configureRouter()
	return r
}

func (rs *restserver) configureRouter() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	rs.router.Use(cors.Handler)
	rs.router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, "pong")
			return
		})
		rapi.Group(func(rclients chi.Router) {
			rclients.Use(rs.checkRole)

			rclients.Group(func(radmin chi.Router) {
				radmin.Use(rs.checkRoleAdmin)
				radmin.Get("/clients/list", rs.clientsList)
				radmin.Get("/psychologist/list", rs.psychologistList)
				radmin.Get("/lesson/list", rs.lessonList)
				radmin.Post("/lessons/pyschologist/{psychologist_id}/client/{client_id}/datetime/{date_time}/set", rs.setLesson)
				radmin.Put("/lesson/{date_time}/psychologist/{psychologist_id}/client/{client_id}/datetime/{new_date_time}/reschedule", rs.rescheduleLesson)
			})
		})
	})
}

//Reschedule an activity. The transfer is possible at any time, including after hours of the psychologist.
func (rs *restserver) rescheduleLesson(w http.ResponseWriter, r *http.Request) {
	psychologistID := chi.URLParam(r, "psychologist_id")
	clientID := chi.URLParam(r, "client_id")

	paramDateTimeOld, err := url.QueryUnescape(chi.URLParam(r, "date_time"))
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	paramDateTimeNew, err := url.QueryUnescape(chi.URLParam(r, "new_date_time"))
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	lessonDatetimeOld, err := time.Parse("2006-01-02 15:04", paramDateTimeOld)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	lessonDatetimeNew, err := time.Parse("2006-01-02 15:04", paramDateTimeNew)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	if err := rs.transportPsychologist.LessonReschedule(psychologistID, clientID, lessonDatetimeOld, lessonDatetimeNew); err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.NoContent(w, r)
}

//Assign a lesson between the client and his psychologist. Recording is possible at any time, including after hours of a psychologist.
func (rs *restserver) setLesson(w http.ResponseWriter, r *http.Request) {
	psychologistID := chi.URLParam(r, "psychologist_id")
	clientID := chi.URLParam(r, "client_id")

	dateTime, err := url.QueryUnescape(chi.URLParam(r, "date_time"))
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	lessonDatetime, err := time.Parse("2006-01-02 15:04", dateTime)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	if err := rs.transportPsychologist.LessonSet(psychologistID, clientID, lessonDatetime); err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.NoContent(w, r)
}

//Get a list of classes: date, name of client, name of psychologist.
func (rs *restserver) lessonList(w http.ResponseWriter, r *http.Request) {
	lessonList, err := rs.transportPsychologist.LessonList()
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	clientsID := make([]*model.Client, 0)
	for _, l := range lessonList {
		clientsID = append(clientsID, &model.Client{ID: l.Client.ID})
	}
	clientsName, err := rs.transportClient.ClientsListByID(clientsID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	for _, l := range lessonList {
		client := findClientName(l.Client.ID, clientsName)
		if client == nil {
			continue
		}
		l.Client.FamilyName = client.FamilyName
		l.Client.Name = client.Name
		l.Client.Patronomic = client.Patronomic
	}
	render.JSON(w, r, lessonList)
}

//Get a list of clients: name of client, name of psychologist, assigned client.
func (rs *restserver) psychologistList(w http.ResponseWriter, r *http.Request) {
	psychologist, err := rs.transportPsychologist.PsychologistList()
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	clientsID := make([]*model.Client, 0)
	for _, p := range psychologist {
		for _, c := range p.Clients {
			clientsID = append(clientsID, &model.Client{ID: c.ID})
		}
	}
	clientsName, err := rs.transportClient.ClientsListByID(clientsID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	for _, p := range psychologist {
		for _, c := range p.Clients {
			client := findClientName(c.ID, clientsName)
			if client == nil {
				continue
			}
			c.FamilyName = client.FamilyName
			c.Name = client.Name
			c.Patronomic = client.Patronomic
		}
	}
	render.JSON(w, r, psychologist)
}

//Get a list of clients: name of client, name of psychologist, assigned client.
func (rs *restserver) clientsList(w http.ResponseWriter, r *http.Request) {
	clientsList, err := rs.transportClient.ClientsList()
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	psychologistID := make([]*model.Psychologist, 0)
	for _, c := range clientsList {
		psychologistID = append(psychologistID, &model.Psychologist{
			ID: c.Psychologist.ID,
		})
	}

	psychologistList, err := rs.transportPsychologist.PsychologistListByID(psychologistID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	for _, c := range clientsList {
		for _, p := range psychologistList {
			if c.Psychologist.ID == p.ID {
				c.Psychologist.FamilyName = p.FamilyName
				c.Psychologist.Name = p.Name
				c.Psychologist.Patronomic = p.Patronomic
			}
		}
	}

	render.JSON(w, r, clientsList)
}

func (rs *restserver) checkRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xrole := r.Header.Get("X-User-Role")
		if strings.TrimSpace(xrole) == "" {
			rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role"))
			return
		}
		for _, ur := range rs.userRoles {
			if ur.name == xrole && ur.isActive {
				next.ServeHTTP(w, r.WithContext(r.Context()))
				return
			}
		}
		rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role"))
	})
}

func (rs *restserver) checkRoleAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xrole := r.Header.Get("X-User-Role")
		if strings.TrimSpace(xrole) != "admin" {
			rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role"))
			return
		}
		for _, ur := range rs.userRoles {
			if ur.name == xrole && ur.isActive {
				next.ServeHTTP(w, r.WithContext(r.Context()))
				return
			}
		}
		rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role"))
	})
}

func findClientName(clientsID string, clietnsName []*model.Client) *model.Client {
	for _, c := range clietnsName {
		if clientsID == c.ID {
			return &model.Client{
				ID:         c.ID,
				FamilyName: c.FamilyName,
				Name:       c.Name,
				Patronomic: c.Patronomic,
			}
		}
	}
	return nil
}

func isTheTime(t time.Time) bool {
	if t.Minute() != 0 {
		return false
	}
	return true
}
