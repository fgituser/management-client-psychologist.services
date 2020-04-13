package transportpsychologist

import (
	"time"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
)

//Transport represents transport between micro services
type Transport interface {
	PsychologistListByID(psychologist []*model.Psychologist) ([]*model.Psychologist, error)
	PsychologistList() ([]*model.Psychologist, error)
	LessonList() ([]*model.Employment, error)
	LessonSet(psychologistID, clientID string, dateTime time.Time) error
	LessonReschedule(psychologistID, clientID string, dateTimeOld, dateTimeNew time.Time) error
}
