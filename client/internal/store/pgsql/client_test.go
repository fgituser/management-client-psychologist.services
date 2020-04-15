package pgsql

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/fgituser/management-client-psychologist.services/client/pkg/database"
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

	rows := sqlmock.NewRows([]string{"psychologist_public_id"}).
		AddRow("75d2cdd6-cf69-44e7-9b28-c47792505d81")
	mock.ExpectQuery(`^select psychologist_public_id from client`).
		WithArgs("48faa486-8e73-4c31-b10f-c7f24c115cda").
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	pID, err := store.PsychologistID("48faa486-8e73-4c31-b10f-c7f24c115cda")
	if err != nil {
		t.Fatal(err)
	}
	if pID == "" {
		t.Fatal("Error psychologistID is empty")
	}
	wantedPsychologistID := "75d2cdd6-cf69-44e7-9b28-c47792505d81"
	assert.Equal(t, wantedPsychologistID, pID)
}

func TestStore_ClientsName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"client_public_id", "family_name", "first_name", "patronymic"}).
		AddRow("48faa486-8e73-4c31-b10f-c7f24c115cda", "Гусев", "Евгений", "Викторович")
	mock.ExpectQuery(`^select c.client_public_id, c.family_name, c.first_name, c.patronymic from clients c`).
		WithArgs("75d2cdd6-cf69-44e7-9b28-c47792505d81a").
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	clientsName, err := store.ClientsName("75d2cdd6-cf69-44e7-9b28-c47792505d81a")
	if err != nil {
		t.Fatal(err)
	}

	wantedCLientsName := []*model.Client{
		{
			ID:           "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName:   "Гусев",
			Name:         "Евгений",
			Patronomic:   "Викторович",
			Psychologist: &model.Psychologist{},
		},
	}
	assert.Equal(t, wantedCLientsName, clientsName)
}

func TestStore_IsAttachment(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(1)
	mock.ExpectQuery(`^select count\(c.id\) from clients c`).
		WithArgs("48faa486-8e73-4c31-b10f-c7f24c115cda", "75d2cdd6-cf69-44e7-9b28-c47792505d81").
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	isAttachment, err := store.IsAttachment("48faa486-8e73-4c31-b10f-c7f24c115cda", "75d2cdd6-cf69-44e7-9b28-c47792505d81")
	if err != nil {
		t.Fatal(err)
	}
	if !isAttachment {
		t.Fatal()
	}
}

func TestStore_ClientsList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"client_public_id", "family_name", "first_name", "patronymic", "psychologist_public_id"}).
		AddRow("48faa486-8e73-4c31-b10f-c7f24c115cda", "Гусев", "Евгений", "Викторович", "75d2cdd6-cf69-44e7-9b28-c47792505d81")
	mock.ExpectQuery(`^select c.client_public_id, c.family_name, c.first_name, c.patronymic, c.psychologist_public_id`).
		WillReturnRows(rows)

	//mock.ExpectCommit()
	store := New(&database.DB{SQL: sqlxDB})
	clientsList, err := store.ClientsList()
	assert.NoError(t, err)
	assert.NotNil(t, clientsList)
	wantedCLientsName := []*model.Client{
		{
			ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
			Psychologist: &model.Psychologist{
				ID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
			},
		},
	}

	assert.Equal(t, wantedCLientsName, clientsList)
}

func TestStore_ClientsNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows([]string{"client_public_id", "family_name", "first_name", "patronymic", "psychologist_public_id"}).
		AddRow("48faa486-8e73-4c31-b10f-c7f24c115cda", "Гусев", "Евгений", "Викторович", "75d2cdd6-cf69-44e7-9b28-c47792505d81")
	mock.ExpectQuery(`^select c.client_public_id, c.family_name, c.first_name, c.patronymic, c.psychologist_public_id`).
		WithArgs(pq.Array([]string{"48faa486-8e73-4c31-b10f-c7f24c115cda"})).
		WillReturnRows(rows)

	store := New(&database.DB{SQL: sqlxDB})
	clientsList, err := store.ClientsNames([]*model.Client{{ID: "48faa486-8e73-4c31-b10f-c7f24c115cda"}})
	assert.NoError(t, err)
	assert.NotNil(t, clientsList)
	wantedCLientsName := []*model.Client{
		{
			ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
			Psychologist: &model.Psychologist{
				ID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
			},
		},
	}
	data, _ := json.Marshal(clientsList)
	fmt.Println(string(data))
	assert.Equal(t, wantedCLientsName, clientsList)
}
