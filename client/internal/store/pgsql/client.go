package pgsql

import (
	"database/sql"
	"strings"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type clients struct {
	ClientPublicID       sql.NullString `db:"client_public_id"`
	FamilyName           sql.NullString `db:"family_name"`
	FirstName            sql.NullString `db:"first_name"`
	Patronomic           sql.NullString `db:"patronymic"`
	PsychologistPublicID sql.NullString `db:"psychologist_public_id"`
}

//ClientsNames get clients name
func (s *Store) ClientsNames(client []*model.Client) ([]*model.Client, error) {
	clientsID := make([]string, 0)
	for _, c := range client {
		clientsID = append(clientsID, c.ID)
	}
	cln := make([]*clients, 0)
	if err := s.db.SQL.Select(&cln, `
		select c.client_public_id, c.family_name, c.first_name, c.patronymic, c.psychologist_public_id 
		from clients c
		where c.client_public_id = any ($1)
	`, pq.Array(clientsID)); err != nil {
		return nil, errors.Wrap(err, "an error accured get clients")
	}
	return convertClientDtoToModelClient(cln), nil
}

//ClientsList get all clients and psychologistID
func (s *Store) ClientsList() ([]*model.Client, error) {

	cDB := make([]*clients, 0)

	err := s.db.SQL.Select(&cDB, `
	select c.client_public_id, c.family_name, c.first_name, c.patronymic, c.psychologist_public_id 
	from clients c `)

	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while get all clients")
	}
	return convertClientDtoToModelClient(cDB), nil
}

//ClientsName get names clients from psychologistID
func (s *Store) ClientsName(psychologistID string) ([]*model.Client, error) {

	if strings.TrimSpace(psychologistID) == "" {
		return nil, errors.New("psychologistID is empty")
	}

	cDB := make([]*clients, 0)

	err := s.db.SQL.Select(&cDB, `
	select c.client_public_id, c.family_name, c.first_name, c.patronymic from clients c 
	where c.psychologist_public_id = $1`, psychologistID)

	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while searching clinets")
	}
	return convertClientDtoToModelClient(cDB), nil
}

//IsAttachment check attachment client from psycholog
func (s *Store) IsAttachment(clientID, psychologistID string) (bool, error) {

	if strings.TrimSpace(clientID) == "" {
		return false, errors.New("clientID is empty")
	}

	if strings.TrimSpace(psychologistID) == "" {
		return false, errors.New("psychologistID is empty")
	}

	var count int64

	err := s.db.SQL.Get(&count, `
	select count(c.id) from clients c
	 where c.client_public_id = $1 and c.psychologist_public_id = $2`, clientID, psychologistID)

	if err != nil {
		return false, errors.Wrap(err, "an error occurred while check attachment client from psychologist")
	}

	if count <= 0 {
		return false, nil
	}
	return true, nil
}

func convertClientDtoToModelClient(clientsDTO []*clients) []*model.Client {
	mclient := make([]*model.Client, 0)
	for _, c := range clientsDTO {
		mclient = append(mclient, &model.Client{
			ID:         c.ClientPublicID.String,
			FamilyName: c.FamilyName.String,
			Name:       c.FirstName.String,
			Patronomic: c.Patronomic.String,
			Psychologist: &model.Psychologist{
				ID: c.PsychologistPublicID.String,
			},
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
