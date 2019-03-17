package shipping

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

var shippingAPIUrl = "https://hei.innovate360.co.uk/api"

// HTTPPoster it's an interface which represents Post method of http.Client
type HTTPPoster interface {
	Post(url, contentType string, body io.Reader) (*http.Response, error)
}

// Client it's a struct which stores APIKey, http client and methods
// which connects you to hey.innovative360.co.uk
type Client struct {
	APIKey string
	HTTPPoster
}

// GetPrices is an end point for getting prices with inventory
func (c *Client) GetPrices(request GetPricesReq) (*GetPricesRes, error) {
	// if true {
	// 	mockResp := &GetPricesRes{}
	// 	GetPricesMock(mockResp)
	// 	return mockResp, nil
	// }
	request.Key = c.APIKey
	bodyReader, w := io.Pipe()
	defer bodyReader.Close()

	ec := make(chan error)
	go func(ec chan error) {
		defer w.Close()
		ec <- json.NewEncoder(w).Encode(request)
	}(ec)

	action := "/quote"
	rawResp, err := c.HTTPPoster.Post(shippingAPIUrl+action, "application/json", bodyReader)
	encodeErr := <-ec
	if encodeErr != nil {
		log.Fatal("Error: Something weird with our struct, it cannot be marshalled to JSON")
		return nil, err
	}
	if err != nil {
		return nil, errors.New("Error occured during request to hey.innovate.360. Details: " + err.Error())
	}
	defer rawResp.Body.Close()
	resp := &GetPricesRes{}
	json.NewDecoder(rawResp.Body).Decode(resp)
	return resp, nil
}
