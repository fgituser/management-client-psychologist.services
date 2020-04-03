package pgsql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
