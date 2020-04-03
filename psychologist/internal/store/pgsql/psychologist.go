package pgsql

import (
	"database/sql"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/users"
	"github.com/pkg/errors"
)

type clients struct {
	ClientPublicID  sql.NullString `db:"client_public_id"`
	EmploeePublicID sql.NullString `db:"employee_public_id"`
}

//FindClients find all clients
func (s *Store) FindClients(employeeID string) ([]*users.Client, error) {

	clients := make([]*clients, 0)

	err := s.db.SQL.Select(&clients, `select client_public_id from clients c
		inner join employee e on e.id = c.employee_id
	where e.employee_public_id = $1`, employeeID)

	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while searching fro clients")
	}
	u := make([]*users.Client, 0)
	for _, c := range clients {
		u = append(u, &users.Client{
			ID: c.ClientPublicID.String,
		})
	}
	return u, nil
}
