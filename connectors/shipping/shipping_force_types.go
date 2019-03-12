package shipping

// GetPricesAddr it's a {from/to} address
type GetPricesAddr struct {
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Country string  `json:"country"`
}

// GetPricesReqItem it's an item in request for moving with its quantity
type GetPricesReqItem struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

// GetPricesReq it's a representation GetPrices Using Inventory Items
type GetPricesReq struct {
	Key         string             `json:"key"`
	Origin      GetPricesAddr      `json:"origin"`
	Destination GetPricesAddr      `json:"destination"`
	Items       []GetPricesReqItem `json:"items"`
	Currency    string             `json:"currency"`
}

type GetPricesResRateProduct struct {
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
type GetPricesResRateSupplier struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
type GetPricesResRateBreakdown struct {
	IncMargin   float64 `json:"inc_margin"`
	IncTax      float64 `json:"inc_tax"`
	TaxEligable int     `json:"tax_eligable"`
}

//GetPricesResRate is a struct for response rate
// TODO: there is a problem with SupplierID, it's 123 and "123" both
type GetPricesResRate struct {
	ID             string                    `json:"id"`
	SupplierID     int                       `json:"-"`
	CurrencySymbol string                    `json:"currency_symbol"`
	CurrencyName   string                    `json:"currency_name"`
	Product        GetPricesResRateProduct   `json:"product"`
	Supplier       GetPricesResRateSupplier  `json:"supplier"`
	BreakDown      GetPricesResRateBreakdown `json:"breakdown"`
}
type GetPricesRes struct {
	Status              string             `json:"status"`
	Description         string             `json:"description"`
	Rates               []GetPricesResRate `json:"rates"`
	Exceptions          []struct{}         `json:"exceptions"`
	CantQuote           bool               `json:"cant_quote"`
	ID                  string             `json:"id"`
	Volume              float64            `json:"volume"`
	VolumeUnit          string             `json:"volume_unit"`
	ConvertedVolume     float64            `json:"converted_volume"`
	ConvertedVolumeUnit string             `json:"converted_volume_unit"`
	Origin              GetPricesAddr      `json:"origin"`
	Destination         GetPricesAddr      `json:"destination"`
}
