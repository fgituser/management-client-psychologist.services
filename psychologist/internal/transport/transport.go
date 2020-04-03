package transport

import "github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"

//Transport represents transport between micro services
type Transport interface {
	GetNamesByID(users []*model.Client) ([]*model.Client, error)
}