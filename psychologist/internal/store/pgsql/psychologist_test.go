package pgsql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/database"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

func TestStore_LessonsListByEmployeeID(t *testing.T) {
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
	ll, err := store.LessonsListByEmployeeID("11e195fc-7010-4e50-8a4d-1d43e9c8e5db")
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

func TestStore_CheckClientAttachment(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(sql.NullInt64{Valid: true, Int64: 1})
	mock.ExpectQuery(`^select count\(c.id\) from clients c`).
		WithArgs("48faa486-8e73-4c31-b10f-c7f24c115cda", "11e195fc-7010-4e50-8a4d-1d43e9c8e5db").
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	isAttachment, err := store.CheckClientAttachment("11e195fc-7010-4e50-8a4d-1d43e9c8e5db", "48faa486-8e73-4c31-b10f-c7f24c115cda")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, isAttachment)
}

func TestStore_LessonIsBusy(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(sql.NullInt64{Valid: true, Int64: 1})
	mock.ExpectQuery(`^select count\(e.id\) from employment e`).
		WithArgs("48faa486-8e73-4c31-b10f-c7f24c115cda",
			time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),      //start_time
			time.Date(2020, 03, 31, 0, 0, 0, 0, time.UTC)). //calendar_id,
		WillReturnRows(rows)

	store := New(&database.DB{SQL: sqlxDB})
	isBusy, err := store.LessonIsBusy("48faa486-8e73-4c31-b10f-c7f24c115cda", time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	//check where datetime is busy
	assert.False(t, isBusy)
}

func TestStore_LessonCanceled(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectExec(`insert into cancellation_employment \(employment_id\)`).
		WithArgs(
			"11e195fc-7010-4e50-8a4d-1d43e9c8e5db",         //employee_id
			time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),      //calendar_id
			time.Date(2020, 03, 31, 0, 0, 0, 0, time.UTC)). //start_time
		WillReturnResult(sqlmock.NewResult(1, 1))

	//mock.ExpectRollback()
	mock.ExpectCommit()

	store := New(&database.DB{SQL: sqlxDB})
	if err := store.LessonCanceled("11e195fc-7010-4e50-8a4d-1d43e9c8e5db",
		time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC)); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStore_EmployeeList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"family_name", "first_name", "patronomic", "employee_public_id"}).
		AddRow("Гусев", "Евгений", "Викторович", "48faa486-8e73-4c31-b10f-c7f24c115cda")
	mock.ExpectQuery(`^select e.family_name, e.first_name, e.patronymic, e.employee_public_id from employee e`).
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	employeeList, err := store.EmployeeList()
	assert.NoError(t, err)
	assert.NotNil(t, employeeList)
	wanted := []*model.Employee{
		{
			ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
		},
	}
	assert.Equal(t, employeeList, wanted)
}

func TestStore_LessonsList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"employee_public_id", "family_name", "first_name", "patronymic", "client_public_id", "calendar_id", "start_time"}).
		AddRow("50faa486-8e73-4c31-b10f-c7f24c115cda", "Гусев", "Евгений", "Викторович",
			"48faa486-8e73-4c31-b10f-c7f24c115cda", time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC), time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC))
	mock.ExpectQuery(`^select em.employee_public_id, em.family_name, em.first_name, em.patronymic, c.client_public_id, s.calendar_id, h.start_time from employment e`).
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	lessonList, err := store.LessonsList()
	assert.NoError(t, err)
	wanted := []*model.Employment{
		{
			Client: &model.Client{
				ID: "48faa486-8e73-4c31-b10f-c7f24c115cda",
			},
			Shedule: []*model.Shedule{
				{
					Employee: &model.Employee{
						ID:         "50faa486-8e73-4c31-b10f-c7f24c115cda",
						FamilyName: "Гусев",
						Name:       "Евгений",
						Patronomic: "Викторович",
					},

					DateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	assert.Equal(t, lessonList, wanted)
}

func TestStore_EmployeesNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"employee_public_id", "family_name", "first_name", "patronymic"}).
		AddRow("50faa486-8e73-4c31-b10f-c7f24c115cda", "Гусев", "Евгений", "Викторович")
	mock.ExpectQuery(`^select e.employee_public_id, e.family_name, e.first_name, e.patronymic from employee e`).
		WithArgs(pq.Array([]string{"50faa486-8e73-4c31-b10f-c7f24c115cda"})).
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	employeeList, err := store.EmployeesNames([]*model.Employee{{ID: "50faa486-8e73-4c31-b10f-c7f24c115cda"}})
	assert.NoError(t, err)
	wanted := []*model.Employee{
		{
			ID:         "50faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
		},
	}
	data, _ := json.Marshal(wanted)
	fmt.Println(string(data))
	assert.Equal(t, employeeList, wanted)
}

func TestStore_LessonsListByEmployeeIDAndClientID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"calendar_id", "start_time"}).
		AddRow(time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC), time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC))
	mock.ExpectQuery(`^select s.calendar_id, h.start_time from employment e`).
		WithArgs("75d2cdd6-cf69-44e7-9b28-c47792505d81", "48faa486-8e73-4c31-b10f-c7f24c115cda").
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	employeeList, err := store.LessonsListByEmployeeIDAndClientID("75d2cdd6-cf69-44e7-9b28-c47792505d81", "48faa486-8e73-4c31-b10f-c7f24c115cda")
	assert.NoError(t, err)
	wanted := []*model.Shedule{
		{
			DateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
		},
	}
	data, _ := json.Marshal(wanted)
	fmt.Println(string(data))
	assert.Equal(t, employeeList, wanted)
}
