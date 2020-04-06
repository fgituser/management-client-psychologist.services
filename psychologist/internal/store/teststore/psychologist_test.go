package teststore

import (
	"reflect"
	"testing"
	"time"

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
			name:    "valid",
			s:       New(),
			args:    args{employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want:    TestClients(t),
			wantErr: false,
		},
		{
			name:    "not valid employyID",
			s:       New(),
			args:    args{employeeID: "85d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty employyID",
			s:       New(),
			args:    args{employeeID: " "},
			want:    nil,
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

func TestStore_LessonsList(t *testing.T) {
	type args struct {
		employeeID string
	}

	tests := []struct {
		name    string
		s       *Store
		args    args
		want    []*model.Employment
		wantErr bool
	}{
		{
			name: "valid",
			s:    New(),
			args: args{employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want: []*model.Employment{
				{
					Client: &model.Client{
						ID: "48faa486-8e73-4c31-b10f-c7f24c115cda",
					},
					Shedule: []*model.Shedule{
						{
							DateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.Local),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "not valid employeeID",
			s:       New(),
			args:    args{employeeID: "90d2cdd6-cf69-44e7-9b28-c47792505d81"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty emploeeID",
			s:       New(),
			args:    args{employeeID: ""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.LessonsList(tt.args.employeeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.LessonsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.LessonsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_SetLesson(t *testing.T) {
	type args struct {
		employeeID string
		clientID   string
		dateTime   time.Time
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			s:    New(),
			args: args{
				employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				clientID:   "48faa486-8e73-4c31-b10f-c7f24c115cda",
				dateTime:   time.Date(2020, 3, 31, 13, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "empty employeeID",
			s:    New(),
			args: args{
				employeeID: "",
				clientID:   "48faa486-8e73-4c31-b10f-c7f24c115cda",
				dateTime:   time.Date(2020, 3, 31, 13, 0, 0, 0, time.Local),
			},
			wantErr: true,
		},
		{
			name: "empty clientID",
			s:    New(),
			args: args{
				employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				clientID:   "",
				dateTime:   time.Date(2020, 3, 31, 13, 0, 0, 0, time.Local),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.SetLesson(tt.args.employeeID, tt.args.clientID, tt.args.dateTime); (err != nil) != tt.wantErr {
				t.Errorf("Store.SetLesson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
