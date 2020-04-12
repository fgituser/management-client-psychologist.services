package transport

import (
	"time"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
)

//Transport represents transport between micro services
type Transport interface {
	ClientLessonList(clientID, psychologistID, userRole string) ([]*model.Shedule, error)
	PsychologistName(psychologistID, userRole string) (*model.Psychologist, error)
	ClientLessonSet(clientID, psychologistID, userRole string, dateTime time.Time) error
	ClientLessonReschedule(clientID, psychologistID, userRole string, dateTimeOld, dateTimeNew time.Time) error
}
