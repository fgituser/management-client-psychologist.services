package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func Test_restserver_clientsList(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/clients/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович","psychologist":{"id":"60faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Себастьянов","name":"Виктор","patronomic":"Андреевич"}}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_psychologistList(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/psychologist/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"id":"60faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Себастьянов","name":"Виктор","patronomic":"Андреевич","clients":[{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"}]}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_lessonList(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("GET", "/api/v1/lesson/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 200)
	expected := []byte(`[{"client":{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"},"shedule":[{"psychologist":{"id":"51faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Соболев","name":"Виктор","patronomic":"Андреевич"},"date_time":"2020-03-31T13:00:00Z"}]}]`)
	assert.Equal(t, bytes.Trim(rr.Body.Bytes(), "\n"), expected)
}

func Test_restserver_setLesson(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("POST", "/api/v1/lessons/pyschologist/80d2cdd6-cf69-44e7-9b28-c47792505d81/client/75d2cdd6-cf69-44e7-9b28-c47792505d81/datetime/2020-03-31%2013%3A00/set", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 204)
}

func Test_restserver_rescheduleLesson(t *testing.T) {
	rest := testRest(t)
	req, err := http.NewRequest("PUT", "/api/v1/lesson/2020-03-31%2013%3A00/psychologist/80d2cdd6-cf69-44e7-9b28-c47792505d81/client/75d2cdd6-cf69-44e7-9b28-c47792505d81/datetime/2020-03-31%2014%3A00/reschedule", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-User-Role", "admin")
	rr := httptest.NewRecorder()
	rest.router.ServeHTTP(rr, req)
	assert.EqualValues(t, rr.Code, 204)
}
