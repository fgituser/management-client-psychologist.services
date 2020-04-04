package service

import (
	"net/http"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/store"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/transport"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func Test_restserver_clientsName(t *testing.T) {
	type fields struct {
		router          chi.Router
		logger          *logrus.Logger
		userRoles       []*UserRole
		store           store.Store
		transportClient transport.Transport
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &restserver{
				router:          tt.fields.router,
				logger:          tt.fields.logger,
				userRoles:       tt.fields.userRoles,
				store:           tt.fields.store,
				transportClient: tt.fields.transportClient,
			}
			rs.clientsName(tt.args.w, tt.args.r)
		})
	}
}
