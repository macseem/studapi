package places

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var googleAPIURL = "https://maps.googleapis.com/maps/api/place"

type httpGetter interface {
	Get(url string) (resp *http.Response, err error)
}

// Client Connector
type Client struct {
	APIKey     string
	HTTPGetter httpGetter
}

// Autocomplete is a method of a Client for getting predictions
func (c *Client) Autocomplete(sessionToken, autocompleteType, input string) (*AutocompleteResponse, error) {
	action := "/autocomplete/json?"
	q := url.Values{}
	q.Set("input", input)
	q.Set("types", autocompleteType)
	q.Set("key", c.APIKey)
	q.Set("sessiontoken", sessionToken)

	// if true {return getCitiesMock(),nil}
	//	log.Print(googleAPIURL + action + q.Encode())
	//return nil, nil
	resp, err := c.HTTPGetter.Get(googleAPIURL + action + q.Encode())
	if err != nil {
		return nil, errors.New("Failed to fetch data. details: " + err.Error())
	}
	defer resp.Body.Close()
	autocompleteResponse := &AutocompleteResponse{}
	err = json.NewDecoder(resp.Body).Decode(autocompleteResponse)
	if err != nil {
		return nil, errors.New("Failed to decode response. details: " + err.Error())
	}
	return autocompleteResponse, err
}

// GetPlaceDetailsByPlaceID retrieves google place details by google place id
func (c *Client) GetPlaceDetailsByPlaceID(sessionToken, id string) (*PlaceDetails, error) {
	action := "/details/json?"
	q := url.Values{}
	q.Set("key", c.APIKey)
	q.Set("sessiontoken", sessionToken)
	q.Set("placeid", id)
	q.Set("fields", "formatted_address,geometry,address_components")
	log.Print(googleAPIURL + action + q.Encode())

	if true {return getPlaceDetailsMock(),nil}
	//log.Print(googleAPIURL + action + q.Encode())
	//return nil, nil
	rawResp, err := c.HTTPGetter.Get(googleAPIURL + action + q.Encode())
	if err != nil {
		return nil, errors.New("Error occured during fetching place details. Error Details: " + err.Error())
	}
	//defer rawResp.Body.Close()
	logRespBytes, _ := ioutil.ReadAll(rawResp.Body)
	rawResp.Body.Close()
	rawResp.Body = ioutil.NopCloser(bytes.NewBuffer(logRespBytes))
	log.Print(string(logRespBytes))
	resp := &GetPlaceDetailsResp{}
	err = json.NewDecoder(rawResp.Body).Decode(resp)
	respBytes, _ := json.Marshal(resp)
	log.Print(string(respBytes))
	if err != nil {
		return nil, errors.New("Cannot decode response. Details: " + err.Error())
	}
	if len(resp.ErrorMessage) > 0 {
		return nil, errors.New("Something went wrong during getting place details. Details: " + resp.ErrorMessage)
	}
	return &resp.Result, nil
}
