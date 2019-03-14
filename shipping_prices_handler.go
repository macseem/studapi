package main

import (
	"clearmove/studapi/connectors/places"
	"clearmove/studapi/connectors/shipping"
	"encoding/json"
	"errors"
	"net/http"
)

// ShippingPricesHandler it's a factory method for the AppHandler of ShippingPrices endpoint
var ShippingPricesHandler AppHandler = func(r *http.Request) *AppRes {
	sr := &ShippingPricesReq{}
	err := json.NewDecoder(r.Body).Decode(sr)
	if err != nil {
		return &AppRes{Code: http.StatusBadRequest, Error: newJSONError("Invalid Request Body", err)}
	}
	getPricesReq, err := mapShippingPricesReqToGetPricesReq(placesClient, sr)
	if err != nil {
		return &AppRes{
			Code:  http.StatusInternalServerError,
			Error: newJSONError("Cannot map our request to the hey.innovative360 request", err),
		}
	}
	resp, err := shippingClient.GetPrices(*getPricesReq)
	if err != nil {
		return &AppRes{
			Code:  http.StatusInternalServerError,
			Error: newJSONError("Error occured during getting prices", err),
		}
	}
	return &AppRes{
		Code: http.StatusOK,
		Data: resp,
	}
}

func mapShippingPricesReqToGetPricesReq(pc *places.Client, sr *ShippingPricesReq) (*shipping.GetPricesReq, error) {
	origin, err := getPricesAddrByPlaceID(pc, sr.SessionToken, sr.From)
	if err != nil {
		return nil, errors.New("Cannot get place details. Details: " + err.Error())
	}
	destination, err := getPricesAddrByPlaceID(pc, sr.SessionToken, sr.To)
	if err != nil {
		return nil, errors.New("Cannot get place details. Details: " + err.Error())
	}

	houseIds := []int{198, 199, 200}
	itemIds := make(map[string]int)
	itemIds["bike"] = 197
	itemIds["box"] = 196
	itemIds["tv"] = 37

	var items []shipping.GetPricesReqItem
	if sr.ItemAmounts != nil {
		items = []shipping.GetPricesReqItem{
			{ID: itemIds["box"], Quantity: sr.ItemAmounts.Box},
			{ID: itemIds["bike"], Quantity: sr.ItemAmounts.Bike},
			{ID: itemIds["tv"], Quantity: sr.ItemAmounts.TV},
		}
	} else {
		items = []shipping.GetPricesReqItem{{ID: houseIds[sr.HouseRoomsAmount], Quantity: 1}}
	}

	request := &shipping.GetPricesReq{
		Origin:      *origin,
		Destination: *destination,
		Items:       items,
		Currency:    "GBP",
	}
	return request, nil
}

func getPricesAddrByPlaceID(pc *places.Client, sessionToken string, id googlePlaceID) (*shipping.Location, error) {
	details, err := pc.GetPlaceDetailsByPlaceID(sessionToken, string(id))
	if err != nil {
		return nil, err
	}
	var countryCode string
	for _, comp := range details.Components {
		if len(countryCode) > 0 {
			break
		}
		for _, cType := range comp.Types {
			if cType == "country" {
				countryCode = comp.ShortName
				break
			}
		}
	}

	return &shipping.Location{
		Address: details.Address,
		Lat:     details.Geometry.Location.Lat,
		Lng:     details.Geometry.Location.Lng,
		Country: countryCode,
	}, nil
}

// ShippingPricesReq is a request struct
// which represents our ShippingPrices endpoint
type ShippingPricesReq struct {
	SessionToken     string        `json:"sessiontoken"`
	From             googlePlaceID `json:"from"`
	To               googlePlaceID `json:"to"`
	ItemAmounts      *itemAmounts  `json:"item_amounts,omitempty"`
	HouseRoomsAmount int           `json:"house_rooms_amount,omitempty"`
}

type itemAmounts struct {
	Box  int `json:"box"`
	Bike int `json:"bike"`
	TV   int `json:"tv"`
}

type googlePlaceID string
