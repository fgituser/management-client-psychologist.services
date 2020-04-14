package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func Test_restserver_clientsName(t *testing.T) {
	rest := testRest(t)

	req, err := http.NewRequest("GET", "/api/v1/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/name", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "psychologist")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Gusev","name":"Evgeniy","patronomic":"Victorovich"}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)

}

func Test_restserver_lessonListByEmployeeID(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/lessons", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "psychologist")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"client":{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Gusev","name":"Evgeniy","patronomic":"Victorovich"},"shedule":[{"date_time":"2020-03-31T13:00:00+07:00"}]}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_lessonSet(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("POST", "/api/v1/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/48faa486-8e73-4c31-b10f-c7f24c115cda/"+
		"lessons/datetime/2020-03-31%2013%3A00/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "psychologist")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
}

func Test_restserver_lessonReschedule(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("PUT", "/api/v1/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/48faa486-8e73-4c31-b10f-c7f24c115cda/"+
		"lessons/datetime/2020-03-31%2013%3A00/reschedule/datetime/2020-03-31%2014%3A00/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "psychologist")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
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

func Test_restserver_employeesList(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/employees/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович","clients":[{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda"}]}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)

}

func Test_restserver_lessonsList(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/lessons/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"client":{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda"},"shedule":[{"employee":{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"},"date_time":"2020-03-31T13:00:00Z"}]}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_employeesListByID(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("POST", "/api/v1/employees/list_by_id", bytes.NewBuffer([]byte(`[{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda"}]`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_employeeNameByID(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/employees/50faa486-8e73-4c31-b10f-c7f24c115cda/name", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "client")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"}`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_lessonByEmployeeIDAndClientID(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/client/48faa486-8e73-4c31-b10f-c7f24c115cda/lessons", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"date_time":"2020-03-31T13:00:00Z"}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_lessonDelete(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("DELETE", "/api/v1/lessons/client/employee/75d2cdd6-cf69-44e7-9b28-c47792505d81/dateteme/2020-03-31%2014%3A00/delete", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 204)
}

func Test_restserver_clientsNameByEmployeeID(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/name", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Gusev","name":"Evgeniy","patronomic":"Victorovich"}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}
