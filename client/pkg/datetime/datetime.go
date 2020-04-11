package datetime

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

//DateTimeJoiner date and time connection
func DateTimeJoiner(d, t sql.NullTime) (time.Time, error) {
	if !d.Valid || !t.Valid {
		return time.Time{}, fmt.Errorf("error dateTimeJoin, not valid date or time date: %v time %v ", d, t)
	}

	sdate := d.Time.Format("2006-01-02")
	stime := t.Time.Format("15:04:05")

	dateTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v %v", sdate, stime))
	if err != nil {
		return time.Time{}, errors.Wrapf(err, " error dateTimeJoin: date:%v time:%v", sdate, stime)
	}
	return dateTime, err
}

//DateTimeSplitUp slpit up datetime to date and time
func DateTimeSplitUp(dateTime *time.Time) (d, t *sql.NullTime, err error) {
	ldate := "2006-01-02"
	ltime := "15:04:05"

	sdate := dateTime.Format(ldate)
	stime := dateTime.Format(ltime)

	dd, err := time.Parse(ldate, sdate)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error split up date")
	}

	tt, err := time.Parse(ltime, stime)
	if err != nil {
		return nil, nil, errors.Wrap(err, "an error accurred while split up time")
	}

	return &sql.NullTime{Valid: true, Time: dd}, &sql.NullTime{Valid: true, Time: tt}, nil
}
