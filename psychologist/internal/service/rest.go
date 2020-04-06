package service

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/store"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/transport"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type restserver struct {
	router          chi.Router
	logger          *logrus.Logger
	userRoles       []*UserRole
	store           store.Store
	transportClient transport.Transport
}

func newRESTServer(router chi.Router, str store.Store, transportClient transport.Transport) *restserver {
	r := &restserver{
		router:          router,
		logger:          logrus.New(),
		store:           str,
		transportClient: transportClient,
	}

	r.configureRouter()
	return r
}

func (rs *restserver) configureRouter() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	rs.router.Use(cors.Handler)
	rs.router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Group(func(remployees chi.Router) {
			remployees.Use(rs.checkoEmploeeID)
			remployees.Use(rs.checkRole)
			remployees.Get("/employees/{employee_id}/clients/name", rs.clientsNameByEmployeeID)
			remployees.Get("/employees/{employee_id}/clients/lessons", rs.lessonListByEmployeeID)
			remployees.Group(func(remployed chi.Router) {
				remployed.Use(rs.checkAttachment)
				remployed.Post("/employess/{employee_id}/clients/{client_id}/lessons/date_time/{date_time}/set", rs.lessonSet)
			})
		})
	})
}

//Schedule an activity with your client. Recording is possible at any time, including non-working
func (rs *restserver) lessonSet(w http.ResponseWriter, r *http.Request) {
	// employeeID := chi.URLParam(r, "employee_id")
	// xrole := r.Header.Get("X-User-Role")
}

//clientsNameByEmployeeID Get a list of your customer names.
func (rs *restserver) clientsNameByEmployeeID(w http.ResponseWriter, r *http.Request) {
	employeeID := chi.URLParam(r, "employee_id")
	xrole := r.Header.Get("X-User-Role")
	clientsID, err := rs.store.FindClients(employeeID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	clientsIDNames, err := rs.transportClient.GetNamesByID(clientsID, employeeID, xrole)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	render.JSON(w, r, clientsIDNames)
}

//lessonListByEmployeeID Get a list of your classes: date, client name
func (rs *restserver) lessonListByEmployeeID(w http.ResponseWriter, r *http.Request) {
	employeeID := chi.URLParam(r, "employee_id")
	xrole := r.Header.Get("X-User-Role")
	ll, err := rs.store.LessonsList(employeeID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	clientsName, err := rs.transportClient.GetNamesByID(employmentToClientID(ll), employeeID, xrole)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, ErrInternal, err)
		return
	}
	for _, l := range ll {
		for _, c := range clientsName {
			if l.Client.ID == c.ID {
				l.Client.Name = c.Name
				l.Client.FamilyName = c.FamilyName
				l.Client.Patronomic = c.Patronomic
				continue
			}
		}
	}
	render.JSON(w, r, ll)
}

func (rs *restserver) checkAttachment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employeeID := chi.URLParam(r, "employee_id")
		clientID := chi.URLParam(r, "client_id")

		isAttachment, err := rs.store.CheckClientAttachment(employeeID, clientID)
		if err != nil {
			rs.sendErrorJSON(w, r, 500, ErrInternal, err)
			return
		}
		if !isAttachment {
			rs.sendErrorJSON(w, r, 400, "clent not attachment to psychologist", nil)
			return
		}

		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func (rs *restserver) lessonIsBusy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employeeID := chi.URLParam(r, "employee_id")
		dateTime := chi.URLParam(r, "date_time")

		dd, err := time.Parse("2006-01-02 15:04:05", dateTime)
		if err != nil {
			rs.sendErrorJSON(w, r, 500, ErrInternal, err)
			return
		}

		isBusy, err := rs.store.LessonIsBusy(employeeID, dd)
		if err != nil {
			rs.sendErrorJSON(w, r, 500, ErrInternal, err)
			return
		}

		if isBusy {
			rs.sendErrorJSON(w, r, 400, "lesson is busy", nil)
			return
		}

		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func (rs *restserver) checkoEmploeeID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employeeID := chi.URLParam(r, "employee_id")
		if strings.TrimSpace(employeeID) == "" {
			rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid employee_id"))
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func (rs *restserver) checkRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xrole := r.Header.Get("X-User-Role")
		if strings.TrimSpace(xrole) == "" {
			rs.sendErrorJSON(w, r, 403, ErrNoAccess, errors.New("not valid employee_id"))
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

func employmentToClientID(lessons []*model.Employment) []*model.Client {
	clientsID := make([]*model.Client, 0)
	for _, l := range lessons {
		clientsID = append(clientsID, &model.Client{ID: l.Client.ID})
	}
	return clientsID
}
