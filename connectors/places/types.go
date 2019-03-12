package places

// GetPlaceDetailsResp it's a struct
// which represents a response of a GetPlaceDetails query
type GetPlaceDetailsResp struct {
	Result       PlaceDetails `json:"result,omitempty"`
	Status       string       `json:"status"`
	ErrorMessage string       `json:"error_message,omitempty"`
}

// PlaceDetails is a struct
// it represents place details response
type PlaceDetails struct {
	PlaceID    string      `json:"place_id,omitempty"`
	Address    string      `json:"formatted_address"`
	Geometry   geometry    `json:"geometry"`
	Components []component `json:"address_components"`
}

type location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type geometry struct {
	Location location `json:"location"`
}

type component struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// AutocompleteResponse it's a response type for a google autocomplete connector
type AutocompleteResponse struct {
	ErrorMessage string       `json:"error_message,omitempty"`
	Predictions  []prediction `json:"predictions"`
	Status       string       `json:"status"`
}

type prediction struct {
	ID                   string `json:"id"`
	Description          string `json:"description"`
	PlaceID              string `json:"place_id"`
	StructuredFormatting struct {
		MainText      string `json:"main_text"`
		SecondaryText string `json:"secondary_text,omitempty"`
	} `json:"structured_formatting"`
}
