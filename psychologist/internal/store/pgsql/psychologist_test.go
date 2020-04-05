package pgsql

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/database"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestPsychologist_FindClients(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"client_public_id"}).
		AddRow("48faa486-8e73-4c31-b10f-c7f24c115cda")
	mock.ExpectQuery(`^select client_public_id from clients`).
		WithArgs("11e195fc-7010-4e50-8a4d-1d43e9c8e5db").
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	uu, err := store.FindClients("11e195fc-7010-4e50-8a4d-1d43e9c8e5db")
	if err != nil {
		t.Fatal(err)
	}
	if len(uu) <= 0 {
		t.Fatal()
	}
	clientID := "48faa486-8e73-4c31-b10f-c7f24c115cda"
	assert.Equal(t, uu[0].ID, clientID)
}

func TestStore_LessonsList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	dd, err := time.Parse("2006-01-02", "2020-03-31")
	if err != nil {
		t.Error(err)
	}
	tt, err := time.Parse("15:04:05", "13:00:00")
	if err != nil {
		t.Error(err)
	}

	rows := sqlmock.NewRows([]string{"client_public_id", "calendar_id", "start_time"}).
		AddRow("48faa486-8e73-4c31-b10f-c7f24c115cda", dd, tt)

	mock.ExpectQuery(`^select c.client_public_id, s.calendar_id, h.start_time from employment`).
		WithArgs("11e195fc-7010-4e50-8a4d-1d43e9c8e5db").
		WillReturnRows(rows)

	store := New(&database.DB{SQL: sqlxDB})
	ll, err := store.LessonsList("11e195fc-7010-4e50-8a4d-1d43e9c8e5db")
	if err != nil {
		t.Error(err)
	}

	expectdDateTime, err := time.Parse("2006-01-02 15:04:05", "2020-03-31 13:00:00")
	if err != nil {
		t.Error(err)
	}
	expectedResult := []*model.Employment{
		{
			Client: &model.Client{
				ID: "48faa486-8e73-4c31-b10f-c7f24c115cda",
			},
			Shedule: []*model.Shedule{
				{
					DateTime: expectdDateTime,
				},
			},
		},
	}
	assert.Equal(t, ll, expectedResult)
}

func Test_dateTimeJoiner(t *testing.T) {
	type args struct {
		d sql.NullTime
		t sql.NullTime
	}

	dd, err := time.Parse("2006-01-02", "2020-03-31")
	if err != nil {
		t.Error(err)
	}
	tt, err := time.Parse("15:04:05", "13:00:00")
	if err != nil {
		t.Error(err)
	}
	expectdDateTime, err := time.Parse("2006-01-02 15:04:05", "2020-03-31 13:00:00")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "valid",
			args: args{d: sql.NullTime{
				Valid: true,
				Time:  dd,
			}, t: sql.NullTime{
				Valid: true,
				Time:  tt,
			}},
			want:    expectdDateTime, //TODO: time.Date
			wantErr: false,
		},
		{
			name: "not valid",
			args: args{d: sql.NullTime{
				Valid: false,
				Time:  dd,
			}, t: sql.NullTime{
				Valid: false,
				Time:  tt,
			}},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dateTimeJoiner(tt.args.d, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("dateTimeJoiner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dateTimeJoiner() = %v, want %v", got, tt.want)
			}
		})
	}
}
