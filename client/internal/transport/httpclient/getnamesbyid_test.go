package httpclient

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient_Do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/psychologist/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/name")
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		assert.Equal(t, req.Header.Get("User-Agent"), "go client")
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, req.Header.Get("X-Role"), "psychologist")
		body, _ := TestResponseGetNamesById(t)
		rw.Write(body)
	}))

	defer server.Close()

	hclient, err := New(server.URL, "go client", server.Client())
	assert.NoError(t, err)

	data, _ := TestSearchingClientsByID(t)
	body, err := hclient.Do(data,
		"/psychologist/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/name",
		"psychologist")
	assert.NoError(t, err)

	b, _ := TestResponseGetNamesById(t)
	assert.NoError(t, err)
	assert.Equal(t, body, b)
}

func Test_decodeGetNameByID(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []*responseGetNamesByID
		wantErr bool
	}{
		{
			name: "valid",
			args: args{data: func() []byte {
				data, _ := TestResponseGetNamesById(t)
				return data
			}()},
			want: func() []*responseGetNamesByID {
				_, data := TestResponseGetNamesById(t)
				return data
			}(),
			wantErr: false,
		},
		{
			name:    "not valid",
			args:    args{data: nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeGetNameByID(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeGetNameByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeGetNameByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_encodeGetNamesByIDToRequest(t *testing.T) {
// 	type args struct {
// 		c []*model.Client
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "valid",
// 			args: args{c: func() []*model.Client {
// 				_, data := TestSearchingClientsByID(t)
// 				return data
// 			}()},
// 			want: func() []byte {
// 				data, _ := TestRequest(t)
// 				return data
// 			}(),
// 			wantErr: false,
// 		},
// 		{
// 			name:    "not valid",
// 			args:    args{c: nil},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := encodeGetNamesByIDToRequest(tt.args.c)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("encodeGetNamesByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("encodeGetNamesByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestHTTPClient_GetNamesByID(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		// assert.Equal(t, req.URL.String(), "/psychologist/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/name")
// 		// assert.Equal(t, req.Header.Get("Accept"), "application/json")
// 		// assert.Equal(t, req.Header.Get("User-Agent"), "go client")
// 		// assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
// 		// assert.Equal(t, req.Header.Get("X-Role"), "psychologist")
// 		body, _ := TestResponseGetNamesById(t)
// 		rw.Write(body)
// 	}))

// 	defer server.Close()

// 	hclient, err := New(server.URL, "go client", server.Client())
// 	assert.NoError(t, err)

// 	_, searchingClientsID := TestSearchingClientsByID(t)

// 	resClientsNames, err := hclient.GetNamesByID(searchingClientsID, "75d2cdd6-cf69-44e7-9b28-c47792505d81", rolePsychologist)
// 	assert.NoError(t, err)

// 	wantClientsNams := TestClients(t)
// 	assert.NoError(t, err)
// 	assert.Equal(t, resClientsNames, wantClientsNams)
// }
