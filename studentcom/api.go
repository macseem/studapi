package studentcom

import (
	"encoding/json"
	"net/url"
)

func GetListAcc(properties GetListAccProperties) (PropertiesResponse, error) {
	resp := PropertiesResponse{}
	uri := NewUri()
	uri.Path = "/properties"
	query := uri.Query()
	query.Set("city_slug", properties.CitySlug)
	query.Set("page_size", properties.PageSize)
	query.Set("country_slug", properties.CountrySlug)
	query.Set("page_number", properties.PageNumber)
	uri.RawQuery = query.Encode()
	body := mockHttpGetList(uri) // TODO: switch to API
	err := json.Unmarshal(body, &resp)
	return resp, err
}

func NewUri() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   "student.com",
	}
}

// TODO
func PostBook(lead NewLead) (response LeadResponse, err error){
	return
}
