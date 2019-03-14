package main

type jsonError struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

// NewJSONErrorResponse creates a response for JSONErrorHandleFunc
func newJSONError(Message string, err error) *jsonError {
	return &jsonError{
		Message: Message,
		Details: err.Error(),
	}
}
