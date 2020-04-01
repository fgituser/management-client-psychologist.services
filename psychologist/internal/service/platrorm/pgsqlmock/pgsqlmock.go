package pgsqlmock

import (
	"github.com/jmoiron/sqlx"
)

// DB ...
type DB struct {
	SQL *sqlx.DB
}

// New ...
func New(psn string, timeout int) (*DB, error) {
	db, err := sqlx.Connect("postgres", psn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{
		SQL: db,
	}, nil
}
