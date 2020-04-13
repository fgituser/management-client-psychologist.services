package service

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/transport"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type restserver struct {
	router                chi.Router
	logger                *logrus.Logger
	userRoles             []*UserRole
	transportPsychologist transport.Transport
}

func newRESTServer(router chi.Router, transportPsychologist transport.Transport) *restserver {
	r := &restserver{
		router:                router,
		logger:                logrus.New(),
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

			})
		})
	})
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

func isTheTime(t time.Time) bool {
	if t.Minute() != 0 {
		return false
	}
	return true
}
