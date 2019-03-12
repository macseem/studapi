package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AppError is a struct for error handler
type AppError struct {
	Error   error
	Message string
	Code    int
}

// AppHandler is a http.Handler interface implementation with error logging
type AppHandler func(http.ResponseWriter, *http.Request) (interface{}, *AppError)

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body.Close() //  must close
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	log.Printf("%s %s %s\n%s", r.RemoteAddr, r.Method, r.URL, bodyBytes)
	resp, appErr := ah(w, r)
	if appErr != nil {
		JSONErrorHandleFunc(w, appErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("ContentType", "application/json")
	byteResponse, err := json.Marshal(resp)
	if err != nil {
		JSONErrorHandleFunc(w, &AppError{Error: err, Message: "Cannot Marshal Response", Code: 500})
	}
	w.Write(byteResponse)
}
