package shipping

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

var ShippingAPIUrl = "https://hei.innovate360.co.uk/api"

type httpPoster interface {
	Post(url, contentType string, body io.Reader) (*http.Response, error)
}

// Client it's a struct which stores APIKey, http client and methods
// which connects you to hey.innovative360.co.uk
type Client struct {
	APIKey     string
	HTTPPoster httpPoster
}

// GetPrices is an end point for getting prices with inventory
func (c *Client) GetPrices(request GetPricesReq) (*GetPricesRes, error) {
	request.Key = c.APIKey
	body, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Error: Something weird with our struct, it cannot be marshalled to JSON")
		return nil, err
	}
	action := "/quote"
	bodyReader := strings.NewReader(string(body))
	rawResp, err := c.HTTPPoster.Post(ShippingAPIUrl+action, "application/json", bodyReader)
	if err != nil {
		return nil, errors.New("Error occured during request to hey.innovate.360. Details: " + err.Error())
	}
	defer rawResp.Body.Close()
	resp := &GetPricesRes{}
	json.NewDecoder(rawResp.Body).Decode(resp)
	return resp, nil
}
