package transportpsychologist

import "github.com/fgituser/management-client-psychologist.services/operator/internal/model"

//Transport represents transport between micro services
type Transport interface {
	PsychologistListByID(psychologist []*model.Psychologist) ([]*model.Psychologist, error)
	PsychologistList() ([]*model.Psychologist, error)
}
