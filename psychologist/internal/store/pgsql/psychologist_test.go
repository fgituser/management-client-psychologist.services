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

	rows := sqlmock.NewRows([]string{"client_public_id", "calendar_id", "start_time"}).
		AddRow("48faa486-8e73-4c31-b10f-c7f24c115cda", time.Date(2020, 03, 31, 13, 0, 0, 0, time.UTC),
			time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC))

	mock.ExpectQuery(`^select c.client_public_id, s.calendar_id, h.start_time from employment`).
		WithArgs("11e195fc-7010-4e50-8a4d-1d43e9c8e5db").
		WillReturnRows(rows)

	store := New(&database.DB{SQL: sqlxDB})
	ll, err := store.LessonsList("11e195fc-7010-4e50-8a4d-1d43e9c8e5db")
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
					DateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
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
				Time:  time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC),
			}, t: sql.NullTime{
				Valid: true,
				Time:  time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),
			}},
			want:    time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "not valid",
			args: args{d: sql.NullTime{
				Valid: false,
				Time:  time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC),
			}, t: sql.NullTime{
				Valid: false,
				Time:  time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),
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

func Test_dateTimeSplitUp(t *testing.T) {
	type args struct {
		dateTime time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantD   sql.NullTime
		wantT   sql.NullTime
		wantErr bool
	}{
		{
			name: "valid",
			args: args{dateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC)},
			wantD: sql.NullTime{
				Valid: true,
				Time:  time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC),
			},
			wantT: sql.NullTime{
				Valid: true,
				Time:  time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, gotT, err := dateTimeSplitUp(&tt.args.dateTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("dateTimeSplitUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, &tt.wantD) {
				t.Errorf("dateTimeSplitUp() gotD = %v, want %v", gotD, tt.wantD)
			}
			if !reflect.DeepEqual(gotT, &tt.wantT) {
				t.Errorf("dateTimeSplitUp() gotT = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
