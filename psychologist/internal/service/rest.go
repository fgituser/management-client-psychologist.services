package service

import (
	"errors"
	"net/http"
	"strings"

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
			remployees.Get("/employees/{employee_id}/clients/name", rs.clientsName)
		})
	})
}

//Get a list of your customer names.
func (rs *restserver) clientsName(w http.ResponseWriter, r *http.Request) {
	employeeID := chi.URLParam(r, "employee_id")
	xrole := r.Header.Get("X-User-Role")
	clientsID, err := rs.store.FindClients(employeeID)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, err)
		return
	}
	clientsIDNames, err := rs.transportClient.GetNamesByID(clientsID, employeeID, xrole)
	if err != nil {
		rs.sendErrorJSON(w, r, 500, err)
	}
	render.JSON(w, r, clientsIDNames)
}

func (rs *restserver) checkoEmploeeID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employeeID := chi.URLParam(r, "employee_id")
		if strings.TrimSpace(employeeID) == "" {
			rs.sendErrorJSON(w, r, 403, errors.New("not valid employee_id"))
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func (rs *restserver) checkRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xrole := r.Header.Get("X-User-Role")
		if strings.TrimSpace(xrole) == "" {
			rs.sendErrorJSON(w, r, 403, errors.New("not valid employee_id"))
			return
		}
		for _, ur := range rs.userRoles {
			if ur.name == xrole && ur.isActive {
				next.ServeHTTP(w, r.WithContext(r.Context()))
				return
			}
		}
		rs.sendErrorJSON(w, r, 403, errors.New("not valid X-User-Role"))
	})
}
