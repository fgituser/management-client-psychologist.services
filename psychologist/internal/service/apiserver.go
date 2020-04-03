package service

import (
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/config"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/store"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/store/pgsql"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/database"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/server"
	"github.com/pkg/errors"
)

//Start the API server
func Start(cfg *config.Configuration) error {
	router := server.New()

	store, err := newDatabase(cfg.DB.DSN)
	if err != nil {
		return errors.Wrap(err, "an error occurred while start api server")
	}
	
	restServer := newRESTServer(router, store)

	server.Start(restServer.router, &server.Config{
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
