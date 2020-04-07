package store

import (
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
)

//Store present store
type Store interface {
	FindClients(employeeID string) ([]*model.Client, error)
	LessonsList(employeeID string) ([]*model.Employment, error)
	SetLesson(employeeID, clientID string, dateTime time.Time) error
	LessonIsBusy(employeeID string, dateTime time.Time) (bool, error)
	LessonCanceled(employeeID string, dateTime time.Time) error
	CheckClientAttachment(employeeID, clientID string) (bool, error)
}
