package service

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
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
				radmin.Get("/clients/list", rs.clientsList)
			})
		})
	})
}

//Get a list of clients: name of client, name of psychologist, assigned client.
func (rs *restserver) clientsList(w http.ResponseWriter, r *http.Request) {
	clientsList, err := rs.transportPsychologist.ClientsList()
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

func isTheTime(t time.Time) bool {
	if t.Minute() != 0 {
		return false
	}
	return true
}
