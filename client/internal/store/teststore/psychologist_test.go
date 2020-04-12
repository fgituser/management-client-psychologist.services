package teststore

import (
	"reflect"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
)

func TestStore_PsychologistID(t *testing.T) {
	type args struct {
		clientID string
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "valid",
			s:       New(),
			args:    args{clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want:    "58faa486-8e73-4c31-b10f-c7f24c115cda",
			wantErr: false,
		},
		{
			name:    "not valid clientID v1",
			s:       New(),
			args:    args{clientID: "58faa486-8e73-4c31-b10f-c7f24c115cda"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "not valid clientID v2",
			s:       New(),
			args:    args{clientID: " "},
			want:    "",
			wantErr: true,
		},
		{
			name:    "not valid clientID v3",
			s:       New(),
			args:    args{clientID: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.PsychologistID(tt.args.clientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.PsychologistID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store.PsychologistID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_ClientsName(t *testing.T) {
	type args struct {
		psychologistID string
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    []*model.Client
		wantErr bool
	}{
		{
			name:    "valid",
			s:       New(),
			args:    args{psychologistID: "75d2cdd6-cf69-44e7-9b28-c47792505d81a"},
			want:    TestClients(t),
			wantErr: false,
		},
		{
			name:    "not valid psycholistID v1",
			s:       New(),
			args:    args{psychologistID: " "},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not valid psycholistID v1",
			s:       New(),
			args:    args{psychologistID: "85d2cdd6-cf69-44e7-9b28-c47792505d81a"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.ClientsName(tt.args.psychologistID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.ClientsName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.ClientsName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_IsAttachment(t *testing.T) {
	type args struct {
		clientID       string
		psychologistID string
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "valid",
			s:       New(),
			args:    args{clientID: "48faa486-8e73-4c31-b10f-c7f24c115cda", psychologistID: "75d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "not valid empty clientID",
			s:       New(),
			args:    args{clientID: "", psychologistID: "75d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want:    false,
			wantErr: true,
		},
		{
			name:    "not valid empty psychologistID",
			s:       New(),
			args:    args{clientID: "48faa486-8e73-4c31-b10f-c7f24c115cda", psychologistID: ""},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.IsAttachment(tt.args.clientID, tt.args.psychologistID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.IsAttachment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store.IsAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_ClientsList(t *testing.T) {
	tests := []struct {
		name    string
		s       *Store
		want    []*model.Client
		wantErr bool
	}{
		{
			name: "valid",
			s:    New(),
			want: []*model.Client{
				{
					ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
					FamilyName: "Гусев",
					Name:       "Евгений",
					Patronomic: "Викторович",
					Psychologist: &model.Psychologist{
						ID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.ClientsList()
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.ClientsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.ClientsList() = %v, want %v", got, tt.want)
			}
		})
	}
}
