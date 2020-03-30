package postgres

import (
	"testing"
)

func Test_New(t *testing.T) {
	db, err := New("postgres://127.0.0.1/netrika?sslmode=disable&user=postgres&password=postgres", 0)
	if err != nil {
		t.Fatal(err)
	}
	var a = 0
	if err := db.Psql.Get(&a, "SELECT 1"); err != nil {
		t.Fatal(err)
	}
	defer db.Psql.Close()
}
