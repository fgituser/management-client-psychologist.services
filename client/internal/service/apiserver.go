package service

import (
	"net/http"

	"github.com/fgituser/management-client-psychologist.services/client/internal/config"
	"github.com/fgituser/management-client-psychologist.services/client/internal/store"
	"github.com/fgituser/management-client-psychologist.services/client/internal/store/pgsql"
	"github.com/fgituser/management-client-psychologist.services/client/internal/transport/httpclient"
	"github.com/fgituser/management-client-psychologist.services/client/pkg/database"
	"github.com/fgituser/management-client-psychologist.services/client/pkg/server"
	"github.com/pkg/errors"
)

//Start the API server
func Start(cfg *config.Configuration) error {
	router := server.New()

	store, err := newDatabase(cfg.DB.DSN)
	if err != nil {
		return errors.Wrap(err, "an error occurred while start api server")
	}

	tranportSvc, err := httpclient.New(cfg.URLServices.ClientsSvcBaseURL, "go client", &http.Client{})
	if err != nil {
		return errors.Wrap(err, "an error accured while start api server")
	}

	restServer := newRESTServer(router, store, tranportSvc)

	server.Start(restServer.router, restServer.logger, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})
	return nil
}

func newDatabase(dsn string) (store.Store, error) {
	db, err := database.New(dsn, 200) //TODO: 200 ?
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while initializing the database")
	}
	return pgsql.New(db), nil
}
