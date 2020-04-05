package pgsql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
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
	fmt.Println(allLessons[0].CalendarID.Time, " ", allLessons[0].StartTime.Time)
	employment, err := lessonsListToEmployment(allLessons)
	if err != nil {
		return nil, errors.Wrap(err, "an error accurred while searching fro clients ")
	}
	return employment, nil
}

//lessonsListToEmployment transformations struct lessonsList on []*model.Employment
func lessonsListToEmployment(allLessons []*lessonsList) ([]*model.Employment, error) {
	e := make([]*model.Employment, 0)
	for _, a := range allLessons {
		shedule := make([]*model.Shedule, 0)
		for _, onelesson := range allLessons {
			if a.ClientPublicID == onelesson.ClientPublicID {

				dateTime, err := dateTimeJoiner(onelesson.CalendarID, onelesson.StartTime)
				if err != nil {
					return nil, errors.Wrap(err, "an error accurred while transformations lessonList to Employment")
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

//dateTimeJoiner date and time connection
func dateTimeJoiner(d, t sql.NullTime) (time.Time, error) {
	if !d.Valid || !t.Valid {
		return time.Time{}, fmt.Errorf("an error accurred while dateTimeJoin, not valid date or time date: %v time %v ", d, t)
	}

	sdate := d.Time.Format("2006-01-02")
	stime := t.Time.Format("15:04:05")

	dateTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v %v", sdate, stime))
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "an error accured with dateTimeJoin: date:%v time:%v", sdate, stime)
	}
	return dateTime, err
}
