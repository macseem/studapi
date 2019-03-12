package main

import (
	"fmt"
	"log"
	"net/http"
)

func errorLog(err *AppError) {
	log.Printf("Error: %d %s\nDetails: %s\n", err.Code, err.Message, err.Error.Error())
}

// JSONErrorHandleFunc is a func which writes to response writer detailed information about an error
func JSONErrorHandleFunc(w http.ResponseWriter, err *AppError) {
	errorLog(err)
	w.WriteHeader(err.Code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{
	"status_code": %d,
	"error": {
		"message": "%s",
		"details": "%s"
	}
}`, err.Code, err.Message, err.Error.Error())))
}
