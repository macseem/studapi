package shipping

// Location it's a {from/to} address
type Location struct {
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Country string  `json:"country"`
}

// GetPricesReqItem It's an item for a request
type GetPricesReqItem struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

// GetPricesReq it's a representation GetPrices Using Inventory Items
type GetPricesReq struct {
	Key         string             `json:"key"`
	Origin      Location      `json:"origin"`
	Destination Location      `json:"destination"`
	Items       []GetPricesReqItem `json:"items"`
	Currency    string             `json:"currency"`
}

// GetPricesRes it's a response of GetPrices api request
type GetPricesRes struct {
	Status              string             `json:"status"`
	Description         string             `json:"description"`
	Rates               []Rate `json:"rates"`
	Exceptions          []struct{}         `json:"exceptions"`
	CantQuote           bool               `json:"cant_quote"`
	ID                  string             `json:"id"`
	Volume              float64            `json:"volume"`
	VolumeUnit          string             `json:"volume_unit"`
	ConvertedVolume     float64            `json:"converted_volume"`
	ConvertedVolumeUnit string             `json:"converted_volume_unit"`
	Origin              Location      `json:"origin"`
	Destination         Location      `json:"destination"`
}

// Rate is a struct for response rate
// TODO: there is a problem with SupplierID, it's 123 and "123" both
type Rate struct {
	ID             string                    `json:"id"`
	SupplierID     int                       `json:"-"`
	CurrencySymbol string                    `json:"currency_symbol"`
	CurrencyName   string                    `json:"currency_name"`
	Product        product   `json:"product"`
	Supplier       supplier  `json:"supplier"`
	BreakDown      breakdown `json:"breakdown"`
}

type product struct {
	ID                  string `json:"id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	DetailedDescription string `json:"detailed_description"`
	ProductType         string `json:"product_type"`
	PackType            string `json:"pack_type"`
	TransitType         string `json:"transit_type"`
	Type                string `json:"type"`
	Icon                string `json:"icon"`
}

type supplier struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type breakdown struct {
	IncMargin   float64 `json:"inc_margin"`
	IncTax      float64 `json:"inc_tax"`
	TaxEligable int     `json:"tax_eligable"`
}
