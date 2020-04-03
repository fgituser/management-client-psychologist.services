package service

import (
	"errors"
	"net/http"
	"strings"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/store"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type restserver struct {
	router chi.Router
	logger *logrus.Logger
	store  store.Store
}

func newRESTServer(router chi.Router, str store.Store) *restserver {
	r := &restserver{
		router: router,
		logger: logrus.New(),
		store:  str,
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
			remployees.Use(checkoEmploeeID)
			remployees.Get("/employees/{employee_id}/clients/name", rs.clientsName)
		})
	})
}

func (rs *restserver) clientsName(w http.ResponseWriter, r *http.Request) {
	employeeID := chi.URLParam(r, "employee_id")
	clints, err := rs.store.FindClients(employeeID)
	if err != nil {
		server.SendErrorJSON(w, r, 500, err)
		return
	}
	render.JSON(w, r, clints)
}

func checkoEmploeeID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employeeID := chi.URLParam(r, "employee_id")
		if strings.TrimSpace(employeeID) == "" {
			server.SendErrorJSON(w, r, 403, errors.New("not valid employee_id"))
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
