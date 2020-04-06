package pgsql

import (
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

func TestStore_SetLesson(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectExec(`^insert into employment \(client_id, shedule_id\)`).
		WithArgs(
			"48faa486-8e73-4c31-b10f-c7f24c115cda",        //client_public_id
			time.Date(2020, 03, 31, 0, 0, 0, 0, time.UTC), //calendar_id
			time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),     //start_time
			"11e195fc-7010-4e50-8a4d-1d43e9c8e5db").       //employee_id
		WillReturnResult(sqlmock.NewResult(1, 1))

	//mock.ExpectRollback()
	mock.ExpectCommit()

	store := New(&database.DB{SQL: sqlxDB})
	if err := store.SetLesson("11e195fc-7010-4e50-8a4d-1d43e9c8e5db",
		"48faa486-8e73-4c31-b10f-c7f24c115cda",
		time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC)); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
