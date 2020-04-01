package service

import (
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/users"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/database"
)

// Psychologist ...
type Psychologist struct {
	db  *database.DB
	dtb DTB
}

//New return new psychologist service
func New(db *postgres.DB, dtb DTB) *Psychologist {
	return &Psychologist{db: db, dtb: dtb}
}

// ClientsName get a list of your customer names.
func (p *Psychologist) ClientsName(employeeID string) ([]*users.Client, error) {

}
