package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB ...
type DB struct {
	Psql *sqlx.DB
}

// New ...
func New(psn string, timeout int) (*DB, error) {
	db, err := sqlx.Connect("postgres", psn)
	if err != nil {
		return nil, err
	}

	if err := db.Select("", "SELCT 1"); err != nil {
		return nil, err
	}
	return &DB{
		Psql: db,
	}, nil
}
