package transport

import "github.com/fgituser/management-client-psychologist.services/operator/internal/model"

//Transport represents transport between micro services
type Transport interface {
	ClientsList() ([]*model.Client, error)
	PsychologistListByID(psychologist []*model.Psychologist) ([]*model.Psychologist, error)
}
