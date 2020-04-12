package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func Test_restserver_clientNames(t *testing.T) {
	rest := testRest(t)

	req, err := http.NewRequest("GET", "/api/v1/client/psychologist/75d2cdd6-cf69-44e7-9b28-c47792505d81a/names", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "psychologist")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	//expected := `[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"},{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Шмельцер","name":"Вячеслав","patronomic":"Николаевич"},{"id":"60faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Виевская","name":"Анастасия","patronomic":"Федоровна"}]`
	assert.NotNil(t, rr.Body)
}

func Test_isTheTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid is the time",
			args: args{time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "valid is not the time",
			args: args{time.Date(2020, 3, 31, 13, 16, 0, 0, time.UTC)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTheTime(tt.args.t); got != tt.want {
				t.Errorf("isTheTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restserver_clientLessons(t *testing.T) {
	rest := testRest(t)

	req, err := http.NewRequest("GET", "/api/v1/client/75d2cdd6-cf69-44e7-9b28-c47792505d81/lessons", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "client")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	//expected := `[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"},{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Шмельцер","name":"Вячеслав","patronomic":"Николаевич"},{"id":"60faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Виевская","name":"Анастасия","patronomic":"Федоровна"}]`
	assert.NotNil(t, rr.Body)
}

func Test_restserver_psychologistName(t *testing.T) {
	rest := testRest(t)

	req, err := http.NewRequest("GET", "/api/v1/client/75d2cdd6-cf69-44e7-9b28-c47792505d81/psychologist/name", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "client")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	//expected := []byte("{\"id\":\"58faa486-8e73-4c31-b10f-c7f24c115cda\",\"family_name\":\"Васкецов\",\"name\":\"Артем\",\"patronomic\":\"Викторович\"}")
	assert.NotNil(t, rr.Body)
	//assert.Equal(t, rr.Body.Bytes(), expected)
}

func Test_restserver_lessonSet(t *testing.T) {
	rest := testRest(t)

	req, err := http.NewRequest("POST", "/api/v1/client/75d2cdd6-cf69-44e7-9b28-c47792505d81/lesson/2020-03-31%2013%3A00/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "client")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 204)
}

func Test_restserver_lessonReschedule(t *testing.T) {
	rest := testRest(t)

	req, err := http.NewRequest("PUT", "/api/v1/client/75d2cdd6-cf69-44e7-9b28-c47792505d81/"+
		"lesson/2020-03-31%2013%3A00/reschedule/datetime/2020-03-31%2014%3A00/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "client")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 204)
}

func Test_restserver_clientsList(t *testing.T) {
	rest := testRest(t)

	req, err := http.NewRequest("GET", "/api/v1/client/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	assert.NotNil(t, rr.Body)
}
