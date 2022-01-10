package response

type SuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
}
