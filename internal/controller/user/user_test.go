package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getData(data map[string]interface{}) []byte {
	body, _ := json.Marshal(data)
	return body
}
func TestShowHandle(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		Id           interface{}
	}{
		{description: "user delete route not recording 11",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "user delete route not recording 3",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 200,
			Id:           "2",
		},
		{description: "user delete route not recording 3",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 200,
			Id:           "3",
		},
		{description: "user delete route wrong id >",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 400,
			Id:           ">",
		},
		{description: "user delete route wrong id 4.5",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "user delete route wrong id true",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "user delete route wrong id *",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 400,
			Id:           "*",
		},
		{description: "user delete route wrong id word",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			expectedCode: 400,
			Id:           "word",
		},
	}
	for _, test := range tests {
		response, _ := http.Get(fmt.Sprintf("%s%s", test.route, test.Id))
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}

}
func TestDestroyHandle(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		ContentType  string
		expectedCode int
		Id           interface{}
	}{
		{description: "user delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "user delete route not recording",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "47",
		},
		{description: "user delete route wrong id true",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "user delete route wrong id 4.5",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "user delete route wrong id -1",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "-1",
		},
	}

	client := http.Client{}

	for _, test := range tests {
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", test.route, test.Id), nil)
		request.Header.Set("Content-Type", test.ContentType)
		response, _ := client.Do(request)
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}
}

func TestUpdateHandle(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		ContentType  string
		expectedCode int
		Id           interface{}
		Data         map[string]interface{}
	}{
		{description: "user update route success process",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 200,
			Id:           "2",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "pass3435f",
				"status":   false,
			}},

		{description: "user update route not recording",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 404,
			Id:           "11",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "pass35t35f",
				"status":   false,
			}},
		{description: "user post route wrong parameter true",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "pasg3g3s",
				"status":   false,
			}},
		{description: "user post route wrong parameter 4.5",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "4.5",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "passg33g3",
				"status":   false,
			}},
		{description: "user post route wrong parameter kelime",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "kelime",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "pass3g33g",
				"status":   false,
			}},
		{description: "user post route wrong parameter >",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           ">",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "passg353g",
				"status":   false,
			}},
		{description: "user post route wrong parameter *",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "*",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "pass3g3g3",
				"status":   false,
			}},
		{description: "user post route wrong parameter *",
			route:        "http://127.0.0.1:3500/api/v1/users/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 422,
			Id:           "3",
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat",
				"password": "passg3g3",
				"status":   false,
			}},
	}
	client := http.Client{}
	for _, test := range tests {
		request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", test.route, test.Id), bytes.NewBuffer(getData(test.Data)))
		request.Header.Set("Content-Type", test.ContentType)
		response, _ := client.Do(request)
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}
}
func TestPostHandle(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		ContentType  string
		expectedCode int
		Data         map[string]interface{}
	}{
		{description: "user post route wrong filed key",
			route:        "http://127.0.0.1:3500/api/v1/users",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"name":      "selo",
				"surname":   "cello",
				"email":     "suat@gmail.com",
				"passwordj": "pass3435f",
				"status":    false,
			}},
		{description: "user post route missing field",
			route:        "http://127.0.0.1:3500/api/v1/users",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"name":      "selo",
				"surname":   "cello",
				"email":     "suat@gmail.com",
				"pasdsword": "pass3435f",
				"status":    false,
			}},
		{description: "user post route missing parameter",
			route:        "http://127.0.0.1:3500/api/v1/users",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"name":     "selo",
				"surname":  "cello",
				"email":    "suat@gmail.com",
				"password": "pass3435f",
			}},
	}

	for _, test := range tests {
		response, _ := http.Post(test.route, test.ContentType, bytes.NewBuffer(getData(test.Data)))
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}
}
