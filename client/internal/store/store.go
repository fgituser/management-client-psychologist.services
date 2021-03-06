package store

import "github.com/fgituser/management-client-psychologist.services/client/internal/model"

//Store presents databases store
type Store interface {
	IsAttachment(clientID, psychologistID string) (bool, error)
	ClientsName(pychologistID string) ([]*model.Client, error)
	PsychologistID(clientID string) (id string, err error)
	ClientsList() ([]*model.Client, error)
	ClientsNames(clients []*model.Client) ([]*model.Client, error)
}

//Image presents image store
type Image interface {
	Upload() error
}
