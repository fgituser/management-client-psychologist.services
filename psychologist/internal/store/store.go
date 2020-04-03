package store

import (
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/users"
)

//Store present store
type Store interface {
	FindClients(employeeID string) ([]*users.Client, error)
}
