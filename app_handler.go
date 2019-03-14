package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// AppRes Common Application Response struct
type AppRes struct {
	Code  int         `json:"status_code"`
	Data  interface{} `json:"data,omitempty"`
	Error *jsonError  `json:"error,omitempty"`
}

// AppHandler is a http.Handler interface implementation with error logging
type AppHandler func(r *http.Request) *AppRes

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")
	bodyTeeBuf := bytes.Buffer{}
	rBody := r.Body
	defer rBody.Close()
	teeBody := ioutil.NopCloser(io.TeeReader(rBody, &bodyTeeBuf))
	r.Body = teeBody
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	resp := ah(r)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	log.Printf("[Body]:\n%s\n[EndBody]", bodyTeeBuf.String())

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
