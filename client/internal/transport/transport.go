package transport

import "github.com/fgituser/management-client-psychologist.services/client/internal/model"

//Transport represents transport between micro services
type Transport interface {
	ClientLessonList(clientID, psychologistID, userRole string) ([]*model.Shedule, error)
}
