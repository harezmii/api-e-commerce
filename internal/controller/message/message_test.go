package message

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
		{description: "message delete route not recording",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "message delete route not recording 17",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			expectedCode: 200,
			Id:           "3",
		},
		{description: "message delete route not recording >",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			expectedCode: 400,
			Id:           ">",
		},
		{description: "message delete route not recording 4.5",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "message delete route not recording true",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "message delete route not recording *",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			expectedCode: 400,
			Id:           "*",
		},
		{description: "message delete route not recording word",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
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
		{description: "message delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "message delete route success case",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "47",
		},
		{description: "message delete route not recording not recording",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "2",
		},
		{description: "message delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "message delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "message delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
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
		{description: "message update route success process",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 200,
			Id:           "3",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email@gmail.com",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
			}},

		{description: "message update route not recording",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 404,
			Id:           "11",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email@gmail.com",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
			}},
		{description: "message post route missing parameter",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email@gmail.com",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
			}},
		{description: "message post route wrong parameter key",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "4.5",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email@gmail.com",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"statusd": true,
			}},
		{description: "message post route email field wrong",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 422,
			Id:           "3",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
			}},
		{description: "message post route wrong parameter kelime",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "kelime",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email@gmail.com",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
			}},
		{description: "message post route wrong parameter >",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           ">",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email@gmail.com",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
			}},
		{description: "message post route wrong parameter *",
			route:        "http://127.0.0.1:3500/api/v1/messages/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "*",
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email@gmail.com",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
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
		{description: "message post route wrong parameter",
			route:        "http://127.0.0.1:3500/api/v1/messages",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"name":    "ndame",
				"email":   "email",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"statuss": true,
			}},
		{description: "message post route unique question name error",
			route:        "http://127.0.0.1:3500/api/v1/messages",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"name":    "name",
				"email":   "email",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
			}},
		{description: "message post route wrong parameter",
			route:        "http://127.0.0.1:3500/api/v1/messages",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"name":    "namewe",
				"email":   "email",
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"statuds": true,
			}},
		{description: "message post route missing parameter",
			route:        "http://127.0.0.1:3500/api/v1/messages",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"phone":   "05468937412",
				"subject": "subject",
				"message": "message",
				"ip":      "178.478.9.6",
				"status":  true,
			}},
	}

	for _, test := range tests {
		response, _ := http.Post(test.route, test.ContentType, bytes.NewBuffer(getData(test.Data)))
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}
}
