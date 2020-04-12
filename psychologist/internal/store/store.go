package store

import (
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
)

//Store present store
type Store interface {
	FindClients(employeeID string) ([]*model.Client, error)
	LessonsListByEmployeeID(employeeID string) ([]*model.Employment, error)
	LessonsList() ([]*model.Employment, error)
	SetLesson(employeeID, clientID string, dateTime time.Time) error
	LessonIsBusy(employeeID string, dateTime time.Time) (bool, error)
	LessonCanceled(employeeID string, dateTime time.Time) error
	CheckClientAttachment(employeeID, clientID string) (bool, error)
	EmployeeList() ([]*model.Employee, error)
	EmployeesNames(employees []*model.Employee) ([]*model.Employee, error)
}
