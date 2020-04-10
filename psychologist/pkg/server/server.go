package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// New instantes new Chi.Router
func New() chi.Router {
	router := chi.NewRouter()
	return router
}

// Config represents server specific config
type Config struct {
	Port                string
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	Debug               bool
}

//Start starts rest api server
func Start(r chi.Router, lg *logrus.Logger, cfg *Config) {
	s := &http.Server{
		Addr:         cfg.Port,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
		Handler:      r,
	}

	go func() {
		lg.Infof("service started from port: %v", cfg.Port)
		if err := s.ListenAndServe(); err != nil {
			lg.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	lg.Info("server stoped")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}
