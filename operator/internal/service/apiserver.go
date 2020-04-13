package service

import (
	"net/http"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/config"
	trclient "github.com/fgituser/management-client-psychologist.services/operator/internal/transportclient/httpclient"
	trpsychologist "github.com/fgituser/management-client-psychologist.services/operator/internal/transportpsychologist/httpclient"
	"github.com/fgituser/management-client-psychologist.services/operator/pkg/server"
	"github.com/pkg/errors"
)

//Start the API server
func Start(cfg *config.Configuration) error {
	router := server.New()

	tranportSvcClient, err := trclient.New(cfg.URLServices.ClientsSvcBaseURL, "go client", &http.Client{})
	if err != nil {
		return errors.Wrap(err, "an error accured while start api server")
	}
	tranportSvcPsychologist, err := trpsychologist.New(cfg.URLServices.PsychologistSvcBaseURL, "go client", &http.Client{})
	if err != nil {
		return errors.Wrap(err, "an error accured while start api server")
	}

	restServer := newRESTServer(router, tranportSvcPsychologist, tranportSvcClient)

	server.Start(restServer.router, restServer.logger, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})
	return nil
}
