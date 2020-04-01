package service

import (
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/employment"
	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/users"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/database"
)

// Service presents psychologist service
type Service interface {
	ClientsName(employeeID string) ([]*users.Client, error)                              //Get a list of your customer names.
	ClientsAppoints(employeeID string) ([]*employment.Employment, error)                 //Get a list of your classes: date, customer name.
	ClientSetAppoint(employeeID string, clientID string) error                           //Schedule an activity with your client. Recording is possible at any time, including non-working
	ClientTransferActivity(employeeID string, clientID string, datatime time.Time) error //Reschedule your occupation. Transfer is possible at any time, including non-working.
}

//DTB presents database repository
type DTB interface {
	FindClients(db *database.DB, employeeID string) ([]*users.Client, error)
	FindAppints(db *database.DB, employeeID string) ([]*employment.Employment, error)
}
