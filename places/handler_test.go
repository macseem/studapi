package places

import (
	"net/http"
	"net/url"
	"testing"
)



func TestAutocompleteHandleFuncMakesCorrectQuery(t *testing.T) {
  expectedResponse:="Response"
	w := MockResponseWriter{
    T: t,
    expectedWrite: []byte(expectedResponse),
  }
	r, _ := http.NewRequest("Get", "http://url.com?query=query&type=cities&sessiontoken=123", nil)
  config := &Config{
    GoogleAPIKey: "apiKey",
  }
	httpClientOld := httpClient
	defer func() { httpClient = httpClientOld }()

  q := url.Values{}
	q.Set("input", "query")
	q.Set("types", "(cities)")
	q.Set("key", "apiKey")
	q.Set("sessiontoken", "123")
  expectedURL:="https://maps.googleapis.com/maps/api/place/autocomplete/json?"+q.Encode()
	httpClient = MockClient{T: t, expectedResponse: expectedResponse, expectedURL:expectedURL}
	AutocompleteHandleFunc(w, r, config)
}
