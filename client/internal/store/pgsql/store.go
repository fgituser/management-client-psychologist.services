package pgsql

import "github.com/fgituser/management-client-psychologist.services/client/pkg/database"

//Store ...
type Store struct {
	db *database.DB
}

//New ...
func New(db *database.DB) *Store {
	return &Store{
		db: db,
	}
}
