package pgsql

import (
	"database/sql"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/datetime"
	"github.com/pkg/errors"
)

type clients struct {
	ClientPublicID  sql.NullString `db:"client_public_id"`
	EmploeePublicID sql.NullString `db:"employee_public_id"`
}

//FindClients find all clients
func (s *Store) FindClients(employeeID string) ([]*model.Client, error) {

	if strings.TrimSpace(employeeID) == "" {
		return nil, errors.New("employeeID is empty")
	}

	clients := make([]*clients, 0)

	err := s.db.SQL.Select(&clients, `
	select client_public_id from clients c
		inner join employee e on e.id = c.employee_id
	where e.employee_public_id = $1`, employeeID)

	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while searching fro clients")
	}
	u := make([]*model.Client, 0)
	for _, c := range clients {
		u = append(u, &model.Client{
			ID: c.ClientPublicID.String,
		})
	}
	return u, nil
}

type lessonsList struct {
	ClientPublicID sql.NullString `db:"client_public_id"`
	CalendarID     sql.NullTime   `db:"calendar_id"`
	StartTime      sql.NullTime   `db:"start_time"`
}

//LessonsList Get a list of your classes: date, client name
func (s *Store) LessonsList(employeeID string) ([]*model.Employment, error) {
	if strings.TrimSpace(employeeID) == "" {
		return nil, errors.New("employeeID is empty")
	}

	allLessons := make([]*lessonsList, 0)

	err := s.db.SQL.Select(&allLessons, `
select c.client_public_id, s.calendar_id, h.start_time from employment e 
		inner join shedule s on s.id  = e.shedule_id
		inner join clients c on c.id = e.client_id
		inner join hours h on h.id  = s.hour_id
		inner join employee em on em.id = s.employee_id 
	where em.employee_public_id = $1 and e.enabled = true`, employeeID)

	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while searching fro clients")
	}

	employment, err := lessonsListToEmployment(allLessons)
	if err != nil {
		return nil, errors.Wrap(err, "an error accurred while searching fro clients ")
	}
	return employment, nil
}

//SetLesson Schedule an activity with your client. Recording is possible at any time, including non-working
func (s *Store) SetLesson(employeeID, clientID string, dateTime time.Time) error {
	if strings.TrimSpace(employeeID) == "" || strings.TrimSpace(clientID) == "" {
		return errors.Errorf("an error accurred while set lesson, empty parametrs employeID:%v clientID:%v", employeeID, clientID)
	}
	dateLesson, timeLesson, err := datetime.DateTimeSplitUp(&dateTime)
	if err != nil {
		return errors.Wrap(err, "an error accurred while set lessons")
	}

	tx := s.db.SQL.MustBegin()
	_, err = tx.Exec(`insert into employment (client_id, shedule_id)
	(
		select c.id client_id, s.id shedule_id from shedule s
			inner join employee e on e.id = s.employee_id 
			inner join hours h on h.id = s.hour_id
			inner join clients c on (c.client_public_id  = $1 and c.employee_id = s.employee_id )
		where s.calendar_id = $2 and h.start_time = $3 and e.employee_public_id = $4
	)`, clientID, dateLesson, timeLesson, employeeID)

	if err != nil {
		return errors.Wrapf(err, "an error accurred while set lesson, empty parametrs employeID:%v clientID:%v", employeeID, clientID)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "an error accurred while set lesson, empty parametrs employeID:%v clientID:%v", employeeID, clientID)
	}
	return nil
}

//lessonsListToEmployment transformations struct lessonsList on []*model.Employment
func lessonsListToEmployment(allLessons []*lessonsList) ([]*model.Employment, error) {
	e := make([]*model.Employment, 0)
	for _, a := range allLessons {
		shedule := make([]*model.Shedule, 0)
		for _, onelesson := range allLessons {
			if a.ClientPublicID == onelesson.ClientPublicID {

				dateTime, err := datetime.DateTimeJoiner(onelesson.CalendarID, onelesson.StartTime)
				if err != nil {
					return nil, errors.Wrap(err, "error transformations lessonList to Employment")
				}

				shedule = append(shedule, &model.Shedule{
					DateTime: dateTime,
				})
			}
			e = append(e, &model.Employment{
				Client:  &model.Client{ID: a.ClientPublicID.String},
				Shedule: shedule,
			})
		}
	}
	return e, nil
}
