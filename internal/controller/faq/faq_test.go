package faq

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetHandle(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{

		{"faq get route", "http://127.0.0.1:3500/api/v1/faqs", 200},
	}
	for _, test := range tests {
		response, _ := http.Get(test.route)
		assert.Equalf(t, test.expectedCode, response.StatusCode, "Faq get handle error")
	}
}
