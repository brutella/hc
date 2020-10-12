package http

// ErrResponse is an error response.
type ErrResponse struct {
	Status int `json:status`
}
