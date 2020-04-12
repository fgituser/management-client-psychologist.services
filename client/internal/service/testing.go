package service

import (
	"testing"

	"github.com/fgituser/management-client-psychologist.services/client/internal/store/teststore"
	"github.com/fgituser/management-client-psychologist.services/client/internal/transport/testtransport"
	"github.com/fgituser/management-client-psychologist.services/client/pkg/server"
	"github.com/sirupsen/logrus"
)

//TestRest rest server
func testRest(t *testing.T) *restserver {

	t.Helper()

	router := server.New()

	rest := restserver{
		router: router,
		logger: logrus.New(),
		userRoles: []*UserRole{
			{
				name:     "psychologist",
				isActive: true,
			}, {
				name: "client",
				isActive: true,
			}, {
				name: "admin",
				isActive: true,
			},
		},
		store:           teststore.New(),
		transportPsychologist: testtransport.New(),
	}
	rest.configureRouter()
	return &rest
}

