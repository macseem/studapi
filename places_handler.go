package main

import (
	"clearmove/studapi/connectors/places"
	"errors"
	"net/http"
)

// NewPlacesHandler is a method to create a handler for places endpoint
func NewPlacesHandler(GoogleAPIKey string) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) (interface{}, *AppError) {
		placesClient := &places.Client{
			APIKey:     GoogleAPIKey,
			HTTPGetter: &http.Client{},
		}
		q := r.URL.Query()
		resp, err := placesClient.Autocomplete(
			q.Get("sessiontoken"), q.Get("type"), q.Get("input"),
		)
		if err != nil {
			return nil, &AppError{
				Message: "Get Error from Google Autocomplete Connector",
				Error:   err,
				Code:    500,
			}
		}
		if len(resp.ErrorMessage) > 0 {
			return nil, &AppError{
				Message: "Get error message from google places API",
				Error:   errors.New(resp.ErrorMessage),
				Code:    500,
			}
		}
		return resp, nil
	}
}
