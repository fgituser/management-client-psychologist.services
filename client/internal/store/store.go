package store

import "github.com/fgituser/management-client-psychologist.services/client/internal/model"

//Store present store
type Store interface {
	IsAttachment(clientID, psychologistID string) (bool, error)
	ClientsName(pychologistID string) ([]*model.Client, error)
	PsychologistID(clientID string) (id string, err error)
}
