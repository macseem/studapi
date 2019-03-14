package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"io"
)

// AppError is a struct for error handler
type AppError struct {
	Error   error
	Message string
	Code    int
}

// AppHandler is a http.Handler interface implementation with error logging
type AppHandler func(*http.Request) (interface{}, *jsonErrorRes)
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")
	bodyTeeBuf := bytes.Buffer{}
	rBody := r.Body
	defer rBody.Close()
	teeBody:=ioutil.NopCloser(io.TeeReader(rBody, &bodyTeeBuf))
	r.Body=teeBody
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	resp, appErr := ah(r)
	log.Printf("[Body]:\n%s\n[EndBody]", bodyTeeBuf.String())
	if appErr != nil {
		JSONErrorHandleFunc(w, appErr)
		return
	}
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		JSONErrorHandleFunc(w, newJSONErrorResponse(500, "Cannot Marshal Response", err.Error()))
	}
}
