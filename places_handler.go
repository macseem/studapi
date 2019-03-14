package main

import (
	"clearmove/studapi/connectors/places"
	"net/http"
)

// NewPlacesHandler is a method to create a handler for places endpoint
func NewPlacesHandler(GoogleAPIKey string) AppHandler {
	return func(r *http.Request) (interface{}, *jsonErrorRes) {
		placesClient := &places.Client{
			APIKey:     GoogleAPIKey,
			HTTPGetter: &http.Client{},
		}
		q := r.URL.Query()
		resp, err := placesClient.Autocomplete(
			q.Get("sessiontoken"), q.Get("type"), q.Get("input"),
		)
		if err != nil {
			return nil, newJSONErrorResponse(http.StatusInternalServerError, "Get Error from Google Autocomplete Connector", err.Error())
		}
		if len(resp.ErrorMessage) > 0 {
			return nil, newJSONErrorResponse(http.StatusInternalServerError, "Get error message from google places API", err.Error())
		}
		return resp, nil
	}
}
