package faq

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetHandle(t *testing.T) {
	tests := []struct {
		description  string
		method       string
		route        string
		expectedCode int
	}{

		{"faq get route", "GET", "http://127.0.0.1:3500/api/v1/faqs", 200},
		{"faq post route", "POST", "http://127.0.0.1:3500/api/v1/faqs", 200},
	}
	for _, test := range tests {
		if test.method == "GET" {
			response, _ := http.Get(test.route)
			assert.Equalf(t, test.expectedCode, response.StatusCode, "Faq get handle error")
		}
		responsePost, _ := http.Post(test.route, "json", nil)
		assert.Equalf(t, test.expectedCode, responsePost.StatusCode, "Faq post handle error")

	}
}
