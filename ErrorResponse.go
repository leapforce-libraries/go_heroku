package heroku

// ErrorResponse stores general API error response
//
type ErrorResponse struct {
	Resource string `json:"resource"`
	ID       string `json:"id"`
	Message  string `json:"message"`
}
