package main

import (
	"net/http"
)

// PlacesHandler is a method to create a handler for places endpoint
var PlacesHandler AppHandler = func(r *http.Request) *AppRes {
	q := r.URL.Query()
	resp, err := placesClient.Autocomplete(
		q.Get("sessiontoken"), q.Get("type"), q.Get("input"),
	)
	if err != nil {
		return &AppRes{
			Code:  http.StatusInternalServerError,
			Error: newJSONError("Get Error from Google Autocomplete Connector", err),
		}
	}
	if len(resp.ErrorMessage) > 0 {
		return &AppRes{
			Code: http.StatusInternalServerError,
			Error: &jsonError{
				Message: "Get error message from google places API",
				Details: resp.ErrorMessage,
			},
		}
	}
	return &AppRes{Code: http.StatusOK, Data: resp.Predictions}
}
