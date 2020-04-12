package service

import (
	"testing"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/store/teststore"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/transport/testtransport"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/server"
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
				name:     "admin",
				isActive: true,
			},
		},
		store:           teststore.New(),
		transportClient: testtransport.New(),
	}
	rest.configureRouter()
	return &rest
}
