package faq

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
		{description: "faq delete route not recording",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "faq delete route not recording 17",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			expectedCode: 200,
			Id:           "17",
		},
		{description: "faq delete route not recording >",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			expectedCode: 400,
			Id:           ">",
		},
		{description: "faq delete route not recording 4.5",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "faq delete route not recording true",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "faq delete route not recording *",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			expectedCode: 400,
			Id:           "*",
		},
		{description: "faq delete route not recording word",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
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
		{description: "faq delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "faq delete route success case",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "14",
		},
		{description: "faq delete route not recording not recording",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "12",
		},
		{description: "faq delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "faq delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "faq delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
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
		{description: "faq update route success process",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 200,
			Id:           "16",
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"answer":   "ne ola",
				"status":   false,
			}},

		{description: "faq update route not recording",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 404,
			Id:           "11",
			Data: map[string]interface{}{
				"question": "test deneme",
				"answer":   "ne ola",
				"status":   false,
			}},
		{description: "faq post route wrong parameter true",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"answer":   "ne ola",
				"status":   false,
			}},
		{description: "faq post route wrong parameter 4.5",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "4.5",
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"answer":   "ne ola",
				"status":   false,
			}},
		{description: "faq post route wrong field",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 422,
			Id:           "17",
			Data: map[string]interface{}{
				"questionn": "test deneme 44",
				"answer":    "ne ola",
				"status":    false,
			}},
		{description: "faq post route missing field",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 422,
			Id:           "17",
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"answer":   "ne ola",
			}},
		{description: "faq post route wrong parameter kelime",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "kelime",
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"answer":   "ne ola",
				"status":   false,
			}},
		{description: "faq post route wrong parameter >",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           ">",
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"answer":   "ne ola",
				"status":   false,
			}},
		{description: "faq post route wrong parameter *",
			route:        "http://127.0.0.1:3500/api/v1/faqs/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "*",
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"answer":   "ne ola",
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
		{description: "faq post route wrong parameter",
			route:        "http://127.0.0.1:3500/api/v1/faqs",
			ContentType:  "application/json",
			expectedCode: 204,
			Data: map[string]interface{}{
				"question": "test deneme 4",
				"answer":   "ne ola",
				"status":   true,
				"deneme":   "d",
			}},
		{description: "faq post route unique question name error",
			route:        "http://127.0.0.1:3500/api/v1/faqs",
			ContentType:  "application/json",
			expectedCode: 204,
			Data: map[string]interface{}{
				"question": "test deneme 478",
				"answer":   "ne ola",
				"status":   true,
			}},
		{description: "faq post route wrong parameter",
			route:        "http://127.0.0.1:3500/api/v1/faqs",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"question": "test deneme 4",
				"answer":   "ne ola",
				"statusd":  true,
			}},
		{description: "faq post route missing parameter",
			route:        "http://127.0.0.1:3500/api/v1/faqs",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"question": "test deneme 44",
				"status":   true,
			}},
	}

	for _, test := range tests {
		response, _ := http.Post(test.route, test.ContentType, bytes.NewBuffer(getData(test.Data)))
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}
}
func TestGetHandle(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{"faq get route offset 0", "http://127.0.0.1:3500/api/v1/faqs?offset=0", 200},
		{"faq get route offset space", "http://127.0.0.1:3500/api/v1/faqs?offset", 200},
		{"faq get route offset bool", "http://127.0.0.1:3500/api/v1/faqs?offset=true", 400},
		{"faq get route offset ?", "http://127.0.0.1:3500/api/v1/faqs?offset=?", 400},
		{"faq get route offset float", "http://127.0.0.1:3500/api/v1/faqs?offset=4.7", 400},
		{"faq get route offset minus", "http://127.0.0.1:3500/api/v1/faqs?offset=-4", 200},
		{"faq get route offset word", "http://127.0.0.1:3500/api/v1/faqs?offset=word", 400},
	}
	for _, test := range tests {
		response, _ := http.Get(test.route)
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)

	}
}
