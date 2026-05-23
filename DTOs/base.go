package dtos

// MessageResponse is used for simple success messages.
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is used for validation and lookup errors.
type ErrorResponse struct {
	Error string `json:"error"`
}
