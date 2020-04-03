package database

import (
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// DB ...
type DB struct {
	SQL *sqlx.DB
}

// New ...
func New(dsn string, timeout int) (*DB, error) {
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("an error occurred while New database: empty dsn")
	}
	db, err := sqlx.Connect("postgres", dsn)
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
