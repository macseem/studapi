package places

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type httpGetter interface {
	Get(url string) (resp *http.Response, err error)
}

var httpClient httpGetter = &http.Client{}

func searchInGoogle(apiKey string, sessionToken string, queryType string, queryString string) (*http.Response, error) {
	// TODO: Add Session, IP check protection
	apiURL := "https://maps.googleapis.com/maps/api/place/autocomplete/json?"
	var q = url.Values{}
	q.Set("input", queryString)
	q.Set("types", "("+queryType+")")
	q.Set("key", apiKey)
	q.Set("sessiontoken", sessionToken)
	return httpClient.Get(apiURL + q.Encode())
}

// AutocompleteHandleFunc is an API Handler for autocomplete
func AutocompleteHandleFunc(w http.ResponseWriter, r *http.Request, config *Config) {
	query := r.URL.Query()
	resp, err := searchInGoogle(config.GoogleAPIKey, query.Get("sessiontoken"), query.Get("type"), query.Get("query"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
