package service

import (
	"net/http"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/config"
	"github.com/fgituser/management-client-psychologist.services/operator/internal/transport/httpclient"
	"github.com/fgituser/management-client-psychologist.services/operator/pkg/server"
	"github.com/pkg/errors"
)

//Start the API server
func Start(cfg *config.Configuration) error {
	router := server.New()

	tranportSvc, err := httpclient.New(cfg.URLServices.ClientsSvcBaseURL, "go client", &http.Client{})
	if err != nil {
		return errors.Wrap(err, "an error accured while start api server")
	}

	restServer := newRESTServer(router, tranportSvc)

	server.Start(restServer.router, restServer.logger, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})
	return nil
}
