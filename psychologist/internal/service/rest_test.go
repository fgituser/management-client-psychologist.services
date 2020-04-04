package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
	//expected := `[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"},{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Шмельцер","name":"Вячеслав","patronomic":"Николаевич"},{"id":"60faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Виевская","name":"Анастасия","patronomic":"Федоровна"}]`
	assert.NotNil(t, rr.Body)

}
