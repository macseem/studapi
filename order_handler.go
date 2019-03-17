package main

import (
	"github.com/macseem/studapi/connectors/shipping"
	"encoding/json"
	"net/http"
	"strconv"
)

type getOrdersRes struct {
	Count  int     `json:"count"`
	Orders []order `json:"orders"`
}

func getOrders(offset, limit int) (*[]order, *AppRes) {
	// TODO: move query to a separate method
	rows, err := db.Query(`select
		id, type, first_name, last_name, email, phone, details,
		extract( epoch from created) as created
		from orders OFFSET $1 LIMIT $2`, offset, limit)
	if err != nil {
		return nil, &AppRes{
			Code:  http.StatusInternalServerError,
			Error: newJSONError("Cannot fetch orders", err),
		}
	}
	defer rows.Close()
	var orders []order
	detailsBuf := ""
	for rows.Next() {
		o := order{}
		err = rows.Scan(&o.ID, &o.Type, &o.FirstName, &o.LastName, &o.Email, &o.Phone, &detailsBuf, &o.Created)
		if err != nil {
			return nil, &AppRes{
				Code:  http.StatusInternalServerError,
				Error: newJSONError("Cannot scan orders", err),
			}
		}
		if err = json.Unmarshal([]byte(detailsBuf), &o.Details); err != nil {
			return nil, &AppRes{
				Code:  http.StatusInternalServerError,
				Error: newJSONError("Cannot Unmarshal details", err),
			}
		}

		// TODO: Implement db.Rows to json Writer instead of appending
		orders = append(orders, o)
	}
	return &orders, nil
}

// getOrdersHandler is a handler for getting a list of orders
var getOrdersHandler = AppHandler(func(r *http.Request) *AppRes {
	var limit, offset int
	var err error
	if offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil {
		return &AppRes{
			Code:  http.StatusBadRequest,
			Error: newJSONError("Invalid Offset", err),
		}
	}
	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil {
		return &AppRes{
			Code:  http.StatusBadRequest,
			Error: newJSONError("Invalid Limit", err),
		}
	}
	orders, appRes := getOrders(offset, limit)
	if appRes != nil {
		return appRes
	}
	gor := getOrdersRes{
		Orders: *orders,
	}
	row := db.QueryRow("select count(id) from orders")
	if err = row.Scan(&gor.Count); err != nil {
		return &AppRes{
			Code:  http.StatusBadRequest,
			Error: newJSONError("Cannot get count of orders", err),
		}
	}

	return &AppRes{Code: http.StatusOK, Data: gor}
})

type order struct {
	ID        int          `json:"id"`
	Status    string       `json:"status"`
	Type      string       `json:"type"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Details   orderDetails `json:"details"`
	Created   float64        `json:"created"`
}
type orderDetails struct {
	ShippingDetails      shippingOrderDetails      `json:"shipping,omitempty"`
	AccommodationDetails accommodationOrderDetails `json:"accommodation,omitempty"`
}
type accommodationOrderDetails struct{}
type shippingOrderDetails struct {
	PricesReq   ShippingPricesReq `json:"pricesRequest"`
	Rates       []shipping.Rate   `json:"rates"`
	Description string            `json:"description"`
	ChozenRate  shipping.Rate     `json:"chozenRate"`
}

var shippingOrderHandler = AppHandler(func(r *http.Request) *AppRes {
	o := &order{}
	err := json.NewDecoder(r.Body).Decode(o)
	if err != nil {
		return &AppRes{Code: http.StatusBadRequest, Error: newJSONError("Cannot Decode Request", err)}
	}
	defer r.Body.Close()
	stringDetails, err := json.Marshal(o.Details)
	if err != nil {
		return &AppRes{Code: http.StatusBadRequest, Error: newJSONError("Cannot Encode Details", err)}
	}
	tx, err := db.Begin()
	if err != nil {
		return &AppRes{Code: http.StatusInternalServerError, Error: newJSONError("Cannot Start Transaction", err)}
	}
	// TODO: move query to a separate method
	var id int
	if err := tx.QueryRow(
		`INSERT INTO orders (type,first_name,last_name,phone,email,details) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`,
		"shipping",
		o.FirstName,
		o.LastName,
		o.Phone,
		o.Email,
		stringDetails,
	).Scan(&id); err != nil {
		tx.Rollback()
		return &AppRes{Code: http.StatusInternalServerError, Error: newJSONError("Cannot Read Id", err)}
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return &AppRes{Code: http.StatusInternalServerError, Error: newJSONError("Cannot Commit Transaction", err)}
	}
	return &AppRes{Code: http.StatusOK, Data: id}
})
