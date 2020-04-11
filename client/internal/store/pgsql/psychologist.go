package pgsql

import (
	"database/sql"
	"strings"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/pkg/errors"
)

type clients struct {
	ClientPublicID sql.NullString `db:"client_public_id"`
	FamilyName     sql.NullString `db:"family_name"`
	FirstName      sql.NullString `db:"first_name"`
	Patronomic     sql.NullString `db:"patronomic"`
}

//ClientsName get names clients from psychologistID
func (s *Store) ClientsName(psychologistID string) ([]*model.Client, error) {

	if strings.TrimSpace(psychologistID) == "" {
		return nil, errors.New("psychologistID is empty")
	}

	cDB := make([]*clients, 0)

	err := s.db.SQL.Select(&cDB, `
	select c.client_public_id, c.family_name, c.first_name, c.patronymic from clients c 
	where c.employee_public_id = $1`, psychologistID)

	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while searching clinets")
	}
	return convertClientDtoToModelClient(cDB), nil
}

func convertClientDtoToModelClient(clientsDTO []*clients) []*model.Client {
	mclient := make([]*model.Client, 0)
	for _, c := range clientsDTO {
		mclient = append(mclient, &model.Client{
			ID:         c.ClientPublicID.String,
			FamilyName: c.FamilyName.String,
			Name:       c.FirstName.String,
			Patronomic: c.Patronomic.String,
		})
	}
	return mclient
}

//PsychologistID find psychologist attachment from client
func (s *Store) PsychologistID(clientID string) (string, error) {

	if strings.TrimSpace(clientID) == "" {
		return "", errors.New("clientID is empty")
	}

	var cID string

	err := s.db.SQL.Get(&cID, `
	select psychologist_public_id from client c
	where client_public_id = $1`, clientID)

	if err != nil {
		return "", errors.Wrap(err, "an error occurred while searching psychologistID from client")
	}
	return cID, nil
}
