package studentcom

type GetListAccProperties struct {
	CitySlug    string `json:"city_slug"`
	PageSize    string `json:"page_size"`
	CountrySlug string `json:"country_slug"`
	PageNumber  string `json:"page_number"`
}

type PropertiesResponse struct {
	Data struct {
		PageInfo struct {
			HasPreviousPage bool `json:"has_previous_page"`
			HasNextPage     bool `json:"has_next_page"`
			NumPages        int  `json:"num_pages"`
		} `json:"page_info"`
		Properties []Properties `json:"properties"`
	} `json:"data"`
}

type Properties struct {
	UnitTypes []struct {
		FloorplanArea   float64 `json:"floorplan_area"`
		BedroomCountMin int     `json:"bedroom_count_min"`
		MaxOccupancy    int     `json:"max_occupancy"`
		KitchenCount    int     `json:"kitchen_count"`
		BedroomCountMax int     `json:"bedroom_count_max"`
		BedCount        int     `json:"bed_count"`
		Listings        []struct {
			AvailableTo   string `json:"available_to"`
			MinDuration   int    `json:"min_duration"`
			DiscountValue string `json:"discount_value"`
			ID            string `json:"id"`
			MoveOut       string `json:"move_out"`
			Availability  string `json:"availability"`
			Type          string `json:"type"`
			MoveIn        string `json:"move_in"`
			PriceMin      string `json:"price_min"`
			DiscountType  string `json:"discount_type"`
			AvailableFrom string `json:"available_from"`
			PriceMax      string `json:"price_max"`
		} `json:"listings"`
		RoomArrangement          string `json:"room_arrangement"`
		Name                     string `json:"name"`
		ID                       string `json:"id"`
		Category                 string `json:"category"`
		BathroomCount            int    `json:"bathroom_count"`
		FloorplanAreaDisplayUnit string `json:"floorplan_area_display_unit"`
	} `json:"unit_types"`
	TotalBeds int `json:"total_beds"`
	Country   struct {
		Name string `json:"name"`
		ID   string `json:"id"`
		Slug string `json:"slug"`
	} `json:"country"`
	Images []struct {
		IsHeroImage bool   `json:"is_hero_image"`
		URL         string `json:"url"`
	} `json:"images"`
	PaymentDepositRules []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"payment_deposit_rules"`
	Address string `json:"address"`
	City    struct {
		Name string `json:"name"`
		ID   string `json:"id"`
		Slug string `json:"slug"`
	} `json:"city"`
	Description  string `json:"description"`
	PostCode     string `json:"post_code"`
	Name         string `json:"name"`
	Currency     string `json:"currency"`
	ID           string `json:"id"`
	BillingCycle string `json:"billing_cycle"`
	Facilities   []struct {
		Name string   `json:"name"`
		Slug string   `json:"slug"`
		Tags []string `json:"tags"`
	} `json:"facilities"`
}

type NewLead struct {
	Email      string `json:"email_address"`
	FirstName  string `json:"first_name"`
	LastName  string `json:"last_name"`
	PropertyID string `json:"property_id"`
}

type LeadResponse struct {
	MoveOutMonth string `json:"move_out_month"`
	ListingID    string `json:"listing_id"`
	Student      struct {
		EmailAddress     string `json:"email_address"`
		LastName         string `json:"last_name"`
		PhoneNumber      string `json:"phone_number"`
		Language         string `json:"language"`
		FirstName        string `json:"first_name"`
		MarketingConsent bool   `json:"marketing_consent"`
		WechatID         string `json:"wechat_id"`
	} `json:"student"`
	BudgetValue               string `json:"budget_value"`
	BudgetFrequency           string `json:"budget_frequency"`
	MoveInMonth               string `json:"move_in_month"`
	DestinationUniversityName string `json:"destination_university_name"`
	BudgetCurrency            string `json:"budget_currency"`
	Tracking                  struct {
		UtmTerm       string `json:"utm_term"`
		Gclid         string `json:"gclid"`
		UtmCampaign   string `json:"utm_campaign"`
		SourceDetails string `json:"source_details"`
		UtmMedium     string `json:"utm_medium"`
		Source        string `json:"source"`
		IPAddress     string `json:"ip_address"`
		SourceDevice  string `json:"source_device"`
		UtmSource     string `json:"utm_source"`
		UtmContent    string `json:"utm_content"`
	} `json:"tracking"`
	PropertyID     string `json:"property_id"`
	PointOfContact struct {
		EmailAddress string `json:"email_address"`
		LastName     string `json:"last_name"`
		PhoneNumber  string `json:"phone_number"`
		FirstName    string `json:"first_name"`
	} `json:"point_of_contact"`
	CityID string `json:"city_id"`
}