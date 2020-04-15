package pgsql

import (
	"database/sql"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/datetime"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type clients struct {
	ClientPublicID  sql.NullString `db:"client_public_id"`
	EmploeePublicID sql.NullString `db:"employee_public_id"`
}

type employee struct {
	FamilyName       sql.NullString `db:"family_name"`
	FirstName        sql.NullString `db:"first_name"`
	Patronymic       sql.NullString `db:"patronymic"`
	EmployeePublicID sql.NullString `db:"employee_public_id"`
	ClientPublicID   sql.NullString `db:"client_public_id"`
}

type employment struct {
	EmploeePublicID   sql.NullString `db:"employee_public_id"`
	EmploeeFamilyName sql.NullString `db:"family_name"`
	EmploeeFirstName  sql.NullString `db:"first_name"`
	EmploeePatronomic sql.NullString `db:"patronymic"`
	ClientPublicID    sql.NullString `db:"client_public_id"`
	CalendarID        sql.NullTime   `db:"calendar_id"`
	StartTime         sql.NullTime   `db:"start_time"`
}

//EmployeeList get all employees
func (s *Store) EmployeeList() ([]*model.Employee, error) {
	empl := make([]*employee, 0)
	if err := s.db.SQL.Select(&empl, `
	select e.family_name, e.first_name, e.patronymic, e.employee_public_id, c.client_public_id from employee e
	left join clients c on e.id = c.employee_id 
	`); err != nil {
		return nil, errors.Wrap(err, "an error accured get all employees")
	}
	return employeeToModelEmployee(empl), nil
}

func employeeToModelEmployee(empl []*employee) []*model.Employee {

	mempl := make([]*model.Employee, 0)

	//we select only unique employees
	employeeName := make(map[string]*employee)
	for _, e := range empl {
		if _, ok := employeeName[e.EmployeePublicID.String]; !ok {
			employeeName[e.EmployeePublicID.String] = e
		}
	}

	//we go through a cycle of unique employees and the internal cycle we collect all the clients of this employee.
	for k, v := range employeeName {
		clients := make([]*model.Client, 0)
		for _, e := range empl {
			if k == e.EmployeePublicID.String {
				clients = append(clients, &model.Client{
					ID: e.ClientPublicID.String,
				})
			}
		}
		mempl = append(mempl, &model.Employee{
			ID:         v.EmployeePublicID.String,
			FamilyName: v.FamilyName.String,
			Name:       v.FirstName.String,
			Patronomic: v.Patronymic.String,
			Clients: func() []*model.Client {
				if clients != nil {
					return clients
				}
				return nil
			}(),
		})
	}

	return mempl
}

//LessonsListByEmployeeIDAndClientID ...
func (s *Store) LessonsListByEmployeeIDAndClientID(employeeID, clientID string) ([]*model.Shedule, error) {
	lessonList := make([]*employment, 0)
	if err := s.db.SQL.Select(&lessonList, `
	select s.calendar_id, h.start_time from employment e 
		inner join sсhedule s on s.id  = e.sсhedule_id 
		inner join clients c on c.id = e.client_id
		inner join hours h on h.id  = s.hour_id
		inner join employee em on em.id = s.employee_id
		left join cancellation_employment ce on ce.employment_id = e.id 
	where em.employee_public_id = $1 and c.client_public_id  = $2 and 
	ce.employment_id is null
	`, employeeID, clientID); err != nil {
		return nil, errors.Wrap(err, "an error accured get employeesNames")
	}
	return employmentToModelSchedule(lessonList)
}

func employmentToModelSchedule(empl []*employment) ([]*model.Shedule, error) {
	schedule := make([]*model.Shedule, 0)
	for _, e := range empl {
		dt, err := datetime.DateTimeJoiner(e.CalendarID, e.StartTime)
		if err != nil {
			return nil, err
		}
		schedule = append(schedule, &model.Shedule{DateTime: dt})
	}
	return schedule, nil
}

//EmployeesNames ...
func (s *Store) EmployeesNames(employees []*model.Employee) ([]*model.Employee, error) {
	employeeID := make([]string, 0)
	for _, e := range employees {
		employeeID = append(employeeID, e.ID)
	}
	empl := make([]*employee, 0)
	if err := s.db.SQL.Select(&empl, `
		select e.employee_public_id, e.family_name, e.first_name, e.patronymic from employee e
		 where e.employee_public_id = any ($1)
	`, pq.Array(employeeID)); err != nil {
		return nil, errors.Wrap(err, "an error accured get employeesNames")
	}
	return employeeToModelEmployee(empl), nil
}

//LessonsList get all lessons
func (s *Store) LessonsList() ([]*model.Employment, error) {
	empl := make([]*employment, 0)
	if err := s.db.SQL.Select(&empl, `
	select em.employee_public_id, em.family_name, em.first_name, em.patronymic, c.client_public_id, s.calendar_id, h.start_time from employment e 
		inner join sсhedule s on s.id  = e.sсhedule_id
		inner join clients c on c.id = e.client_id
		inner join hours h on h.id  = s.hour_id
		inner join employee em on em.id = s.employee_id
		left join cancellation_employment ce on ce.employment_id = e.id 
	where ce.employment_id is null
	`); err != nil {
		return nil, errors.Wrap(err, "an error accured get all lessons")
	}
	return employmentToModelEmployment(empl)
}

func employmentToModelEmployment(empl []*employment) ([]*model.Employment, error) {
	mempl := make([]*model.Employment, 0)
	for _, e := range empl {
		mempl = append(mempl, &model.Employment{
			Client: &model.Client{
				ID: e.ClientPublicID.String,
			},
		})
	}
	for _, m := range mempl {
		for _, e := range empl {
			if m.Client.ID == e.ClientPublicID.String {
				dt, err := datetime.DateTimeJoiner(e.CalendarID, e.StartTime)
				if err != nil {
					return nil, err
				}
				m.Shedule = append(m.Shedule, &model.Shedule{
					Employee: &model.Employee{
						ID:         e.EmploeePublicID.String,
						FamilyName: e.EmploeeFamilyName.String,
						Name:       e.EmploeeFirstName.String,
						Patronomic: e.EmploeePatronomic.String,
					},
					DateTime: dt,
				})
			}
		}
	}
	return mempl, nil
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

//LessonsListByEmployeeID Get a list of your classes: date, client name
func (s *Store) LessonsListByEmployeeID(employeeID string) ([]*model.Employment, error) {
	if strings.TrimSpace(employeeID) == "" {
		return nil, errors.New("employeeID is empty")
	}

	allLessons := make([]*lessonsList, 0)

	err := s.db.SQL.Select(&allLessons, `
	select c.client_public_id, s.calendar_id, h.start_time from employment e 
		inner join sсhedule s on s.id  = e.sсhedule_id
		inner join clients c on c.id = e.client_id
		inner join hours h on h.id  = s.hour_id
		inner join employee em on em.id = s.employee_id
		left join cancellation_employment ce on ce.employment_id = e.id 
	where em.employee_public_id = $1 and 
	ce.employment_id is null`, employeeID)

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
	_, err = tx.Exec(`insert into employment (client_id, schedule_id)
	(
		select c.id client_id, s.id schedule_id from schedule s
			inner join employee e on e.id = s.employee_id 
			inner join hours h on h.id = s.hour_id
			inner join clients c on (c.client_public_id  = $1 and c.employee_id = s.employee_id ) //TODO: change where
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

//CheckClientAttachment ...
func (s *Store) CheckClientAttachment(employeeID, clientID string) (bool, error) {
	if strings.TrimSpace(employeeID) == "" || strings.TrimSpace(clientID) == "" {
		return false, errors.Errorf("an error accured check client attachment: not valid parametrs employeeID:%v clientID:%v", employeeID, clientID)
	}

	var count sql.NullInt64

	err := s.db.SQL.Get(&count, `
	select count(c.id) from clients c
	inner join employee e on e.id = c.employee_id
	where c.client_public_id = $1 and e.employee_public_id = $2`, clientID, employeeID)

	if err != nil {
		return false, errors.Wrap(err, "an error accurred while check client attachment")
	}

	if count.Int64 <= 0 {
		//client not attachment to psychologist
		return false, nil
	}
	return true, nil
}

//LessonIsBusy ...
func (s *Store) LessonIsBusy(employeeID string, dateTime time.Time) (bool, error) {
	if strings.TrimSpace(employeeID) == "" {
		return false, errors.Errorf("an error accured while check lesson free datetime: not valid parametrs employeeID:%v", employeeID)
	}

	dateLesson, timeLesson, err := datetime.DateTimeSplitUp(&dateTime)
	if err != nil {
		return false, errors.Wrap(err, "an error accurred while check lesson is busy: datetime split up")
	}

	var count sql.NullInt64

	err = s.db.SQL.Get(&count, `
	select count(e.id) from employment e 
		inner join schedule s on s.id = e.schedule_id
		inner join hours h on h.id = s.hour_id
		inner join employee empl on empl.id = s.employee_id
		left join cancellation_employment ce on ce.employment_id = e.id 
	where empl.employee_public_id = $1' and 
		h.start_time = $2 and s.calendar_id = $3 and 
		ce.employment_id is null`, employeeID, timeLesson, dateLesson)

	if err != nil {
		return false, errors.Wrap(err, "an error accurrd while chekc lesson is busy: Get")
	}

	if count.Int64 == 0 {
		//datetime free
		return true, nil
	}
	//datetime busy
	return false, nil
}

//LessonCanceled canceled lesson
func (s *Store) LessonCanceled(employeeID string, dateTime time.Time) error {
	if strings.TrimSpace(employeeID) == "" {
		return errors.Errorf("an error accurred while caneled lesson, empty parametrs employeID:%v", employeeID)
	}
	dateLesson, timeLesson, err := datetime.DateTimeSplitUp(&dateTime)
	if err != nil {
		return errors.Wrap(err, "an error accurred while canceled lesson")
	}

	tx := s.db.SQL.MustBegin()
	_, err = tx.Exec(`
	insert into cancellation_employment (employment_id)
	(
		select e.id from employment e 
			inner join schedule s on s.id = e.schedule_id
			inner join hours h on h.id = s.hour_id
			inner join employee empl on empl.id = s.employee_id
			left join cancellation_employment ce on ce.employment_id = e.id 
		where empl.employee_public_id = $1 and 
			h.start_time = $2 and s.calendar_id = $3 and
			ce.employment_id is null
	)`, employeeID, timeLesson, dateLesson)

	if err != nil {
		return errors.Errorf("an error accurred while caneled lesson, empty parametrs employeID:%v", employeeID)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Errorf("an error accurred while caneled lesson, empty parametrs employeID:%v", employeeID)
	}
	return nil
}

//lessonsListToEmployment transformations struct lessonsList on []*model.Employment
func lessonsListToEmployment(allLessons []*lessonsList) ([]*model.Employment, error) {
	e := make([]*model.Employment, 0)
	for _, a := range allLessons {
		schedule := make([]*model.Shedule, 0)

		dateTime, err := datetime.DateTimeJoiner(a.CalendarID, a.StartTime)
		if err != nil {
			return nil, errors.Wrap(err, "error transformations lessonList to Employment")
		}

		schedule = append(schedule, &model.Shedule{
			DateTime: dateTime,
		})

		e = append(e, &model.Employment{
			Client:  &model.Client{ID: a.ClientPublicID.String},
			Shedule: schedule,
		})

	}
	return e, nil
}
