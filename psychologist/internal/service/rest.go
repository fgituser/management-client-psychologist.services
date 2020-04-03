package service

import (
	"net/http"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/store"
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
	rs.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "pong")
		return
	})
}
