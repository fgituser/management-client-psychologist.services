package teststore

import (
	"reflect"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
)

func TestStore_FindClients(t *testing.T) {
	type args struct {
		employeeID string
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    []*model.Client
		wantErr bool
	}{
		{
			name: "valid",
			s: New(),
			args: args{employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want: TestClients(t),
			wantErr: false,
		},
		{
			name: "not valid employyID",
			s: New(),
			args: args{employeeID: "85d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want: nil,
			wantErr: true,
		},
		{
			name: "empty employyID",
			s: New(),
			args: args{employeeID: " "},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.FindClients(tt.args.employeeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.FindClients() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.FindClients() = %v, want %v", got, tt.want)
			}
		})
	}
}
