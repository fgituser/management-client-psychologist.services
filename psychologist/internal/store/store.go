package store

import "github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"

//Store present store
type Store interface {
	FindClients(employeeID string) ([]*model.Client, error)
}
