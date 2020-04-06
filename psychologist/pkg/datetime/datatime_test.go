package datetime

import (
	"database/sql"
	"reflect"
	"testing"
	"time"
)

func Test_dateTimeJoiner(t *testing.T) {
	type args struct {
		d sql.NullTime
		t sql.NullTime
	}

	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "valid",
			args: args{d: sql.NullTime{
				Valid: true,
				Time:  time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC),
			}, t: sql.NullTime{
				Valid: true,
				Time:  time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),
			}},
			want:    time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "not valid",
			args: args{d: sql.NullTime{
				Valid: false,
				Time:  time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC),
			}, t: sql.NullTime{
				Valid: false,
				Time:  time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),
			}},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateTimeJoiner(tt.args.d, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("dateTimeJoiner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dateTimeJoiner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dateTimeSplitUp(t *testing.T) {
	type args struct {
		dateTime time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantD   sql.NullTime
		wantT   sql.NullTime
		wantErr bool
	}{
		{
			name: "valid",
			args: args{dateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC)},
			wantD: sql.NullTime{
				Valid: true,
				Time:  time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC),
			},
			wantT: sql.NullTime{
				Valid: true,
				Time:  time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, gotT, err := DateTimeSplitUp(&tt.args.dateTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("dateTimeSplitUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, &tt.wantD) {
				t.Errorf("dateTimeSplitUp() gotD = %v, want %v", gotD, tt.wantD)
			}
			if !reflect.DeepEqual(gotT, &tt.wantT) {
				t.Errorf("dateTimeSplitUp() gotT = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
