package service

import (
	"testing"

	trclient "github.com/fgituser/management-client-psychologist.services/operator/internal/transportclient/testtransport"
	trpsychologist "github.com/fgituser/management-client-psychologist.services/operator/internal/transportpsychologist/testtransport"
	"github.com/fgituser/management-client-psychologist.services/operator/pkg/server"
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
				name:     "client",
				isActive: true,
			}, {
				name:     "admin",
				isActive: true,
			},
		},
		transportPsychologist: trpsychologist.New(),
		transportClient:       trclient.New(),
	}
	rest.configureRouter()
	return &rest
}
