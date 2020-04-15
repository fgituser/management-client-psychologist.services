package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/fgituser/management-client-psychologist.services/client/internal/store"
	"github.com/fgituser/management-client-psychologist.services/client/internal/transport"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type restserver struct {
	router                chi.Router
	logger                *logrus.Logger
	userRoles             []*userRole
	store                 store.Store
	transportPsychologist transport.Transport
}

func newRESTServer(router chi.Router, str store.Store, userRoles []*userRole, transportPsychologist transport.Transport) *restserver {
	r := &restserver{
		router:                router,
		logger:                logrus.New(),
		store:                 str,
		userRoles:             userRoles,
		transportPsychologist: transportPsychologist,
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
	metrics := chiprometheus.NewMiddleware("client-service")
	rs.router.Use(metrics)
	rs.router.Use(cors.Handler)
	rs.router.Handle("/metrics", promhttp.Handler())
	rs.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "pong")
		return
	})
	rs.router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Group(func(rclients chi.Router) {
			rclients.Use(rs.checkRole)

			rclients.Get("/client/{client_id}/lessons", rs.clientLessons)
			rclients.Get("/client/psychologist/{psychologist_id}/name", rs.clientNames)
			rclients.Get("/client/{client_id}/psychologist/name", rs.psychologistName)
			rclients.Post("/client/{client_id}/lesson/{date_time}/set", rs.lessonSet)
			rclients.Put("/client/{client_id}/lesson/{date_time}/reschedule/datetime/{new_date_time}/set", rs.lessonReschedule)

			rclients.Group(func(radmin chi.Router) {
				radmin.Use(rs.checkRoleAdmin)
				radmin.Get("/client/list", rs.clientsList)
				radmin.Post("/client/list_by_id", rs.clientsListByID)
			})
		})
	})
}

// get clients list by id
func (rs *restserver) clientsListByID(w http.ResponseWriter, r *http.Request) {
	var req []struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	clientsID := make([]*model.Client, 0)
	for _, rq := range req {
		clientsID = append(clientsID, &model.Client{ID: rq.ID})
	}

	clietns, err := rs.store.ClientsNames(clientsID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.JSON(w, r, clietns)

}

//Get clients list
func (rs *restserver) clientsList(w http.ResponseWriter, r *http.Request) {
	clientList, err := rs.store.ClientsList()
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.JSON(w, r, clientList)
}

//Reschedule your occupation. Time transfer is possible only for the working time of the psychologist.
func (rs *restserver) lessonReschedule(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "client_id")
	userRole := r.Header.Get("X-User-Role")
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

	if !isTheTime(lessonDatetimeNew) {
		rs.sendErrorJSON(w, r, 400, "a lesson can only be scheduled at the beginning of the hour", nil)
		return
	}

	psychologistID, err := rs.store.PsychologistID(clientID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	if err := rs.transportPsychologist.ClientLessonReschedule(clientID, psychologistID, userRole, lessonDatetimeOld, lessonDatetimeNew); err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.NoContent(w, r)
}

//Sign up for a lesson to a psychologist. Recording is possible only for the working time of the psychologist
func (rs *restserver) lessonSet(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "client_id")
	userRole := r.Header.Get("X-User-Role")

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

	if !isTheTime(lessonDatetime) {
		rs.sendErrorJSON(w, r, 400, "a lesson can only be scheduled at the beginning of the hour", nil)
		return
	}

	psychologistID, err := rs.store.PsychologistID(clientID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	if err := rs.transportPsychologist.ClientLessonSet(clientID, psychologistID, userRole, lessonDatetime); err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}

	render.NoContent(w, r)
}

//Get the name of your psychologist.
func (rs *restserver) psychologistName(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "client_id")
	userRole := r.Header.Get("X-User-Role")

	psychologistID, err := rs.store.PsychologistID(clientID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	psychologistName, err := rs.transportPsychologist.PsychologistName(psychologistID, userRole)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.JSON(w, r, psychologistName)
}

//get clients name from psychologist id
func (rs *restserver) clientNames(w http.ResponseWriter, r *http.Request) {
	psychologistID := chi.URLParam(r, "psychologist_id")
	clientsNames, err := rs.store.ClientsName(psychologistID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.JSON(w, r, clientsNames)
}

//Get a list of your classes
func (rs *restserver) clientLessons(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "client_id")
	userRole := r.Header.Get("X-User-Role")

	psychologistID, err := rs.store.PsychologistID(clientID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	lessonList, err := rs.transportPsychologist.ClientLessonList(clientID, psychologistID, userRole)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.JSON(w, r, lessonList)
}

func (rs *restserver) checkAttachment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := chi.URLParam(r, "client_id")
		psychologistID := chi.URLParam(r, "psychologist_id")
		if strings.TrimSpace(clientID) == "" {
			rs.sendErrorJSON(w, r, 400, "not valid psychologist_id", nil) //TODO: status warn stdout
			return
		}
		if strings.TrimSpace(psychologistID) == "" {
			rs.sendErrorJSON(w, r, 400, "not valid clinetID", nil) //TODO: status warn stdout
			return
		}
		isAttachment, err := rs.store.IsAttachment(clientID, psychologistID)
		if err != nil {
			rs.sendErrorJSON(w, r, 500, ErrInternal, err)
			return
		}

		if !isAttachment {
			rs.sendErrorJSON(w, r, 400, ErrNoAttachment, nil)
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func (rs *restserver) checkRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xrole := r.Header.Get("X-User-Role")
		if strings.TrimSpace(xrole) == "" {
			rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role - "+xrole))
			return
		}
		for _, ur := range rs.userRoles {
			if ur.name == xrole && ur.isActive {
				next.ServeHTTP(w, r.WithContext(r.Context()))
				return
			}
		}
		rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role - "+xrole))
	})
}

func (rs *restserver) checkRoleAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xrole := r.Header.Get("X-User-Role")
		if strings.TrimSpace(xrole) != "admin" {
			rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role - "+xrole))
			return
		}
		for _, ur := range rs.userRoles {
			if ur.name == xrole && ur.isActive {
				next.ServeHTTP(w, r.WithContext(r.Context()))
				return
			}
		}
		rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid X-User-Role - "+xrole))
	})
}

func isTheTime(t time.Time) bool {
	if t.Minute() != 0 {
		return false
	}
	return true
}
