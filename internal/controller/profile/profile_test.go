package profile

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
		{description: "profile delete route not recording 11",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "profile delete route not recording 3",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			expectedCode: 404,
			Id:           "3",
		},
		{description: "profile delete route wrong id >",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			expectedCode: 400,
			Id:           ">",
		},
		{description: "profile delete route wrong id 4.5",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "profile delete route wrong id true",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "profile delete route wrong id *",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			expectedCode: 400,
			Id:           "*",
		},
		{description: "profile delete route wrong id word",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
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
		{description: "profile delete route not recording ",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "11",
		},
		{description: "profile delete route not recording",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json",
			expectedCode: 404,
			Id:           "47",
		},
		{description: "profile delete route wrong id true",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
		},
		{description: "profile delete route wrong id 4.5",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json",
			expectedCode: 400,
			Id:           "4.5",
		},
		{description: "profile delete route wrong id -1",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
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
		{description: "profile update route success process",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 200,
			Id:           "1",
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
			}},

		{description: "profile update route not recording",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 404,
			Id:           "11",
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
			}},
		{description: "profile post route wrong parameter true",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "true",
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
			}},
		{description: "profile post route wrong parameter 4.5",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "4.5",
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
			}},
		{description: "profile post route wrong parameter kelime",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "kelime",
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
			}},
		{description: "profile post route wrong parameter >",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           ">",
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
			}},
		{description: "profile post route wrong parameter *",
			route:        "http://127.0.0.1:3500/api/v1/profiles/",
			ContentType:  "application/json; charset=utf-8",
			expectedCode: 400,
			Id:           "*",
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
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
		{description: "profile post route wrong filed key",
			route:        "http://127.0.0.1:3500/api/v1/profiles",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"userId4": 2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
				"phone":   "03124562678",
			}},
		{description: "profile post route missing field",
			route:        "http://127.0.0.1:3500/api/v1/profiles",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"userId":  2,
				"image":   "resim.jpg",
				"address": "gokceyurt",
			}},
		{description: "profile post route missing parameter",
			route:        "http://127.0.0.1:3500/api/v1/profiles",
			ContentType:  "application/json",
			expectedCode: 422,
			Data: map[string]interface{}{
				"userId":   2,
				"image":    "resim.jpg",
				"addressf": "gokceyurt",
				"phone":    "03124562678",
			}},
	}

	for _, test := range tests {
		response, _ := http.Post(test.route, test.ContentType, bytes.NewBuffer(getData(test.Data)))
		assert.Equalf(t, test.expectedCode, response.StatusCode, test.description)
	}
}
