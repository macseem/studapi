package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type jsonError struct {
	Message string `json:"message"`
	Details string `json:"details"`
}
type jsonErrorRes struct {
	Code int `json:"status_code"`
	Error jsonError `json:"error"`
}
func (j *jsonErrorRes) String() string {
	return fmt.Sprintf("Error: {Code: [%d] Message: [%s] Details: [%s]}", j.Code, j.Error.Message, j.Error.Details)
}
// NewJSONErrorResponse creates a response for JSONErrorHandleFunc
func newJSONErrorResponse(Code int, Message, Details string) *jsonErrorRes {
	return &jsonErrorRes{
		Code,
		jsonError {
			Message,
			Details,
		},
	}
}

// JSONErrorHandleFunc is a func which writes to response writer detailed information about an error
func JSONErrorHandleFunc(w http.ResponseWriter, err *jsonErrorRes) {
	log.Fatal(err)
	w.WriteHeader(err.Code)
	writeErr := json.NewEncoder(w).Encode(err)
	log.Fatal(writeErr)
}
