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
