package store

import "github.com/fgituser/management-client-psychologist.services/client/internal/model"

//Store present store
type Store interface {
	ClientsName(pychologistID string) ([]*model.Client, error)
	PsychologistID(clientID string) (id string, err error)
}
