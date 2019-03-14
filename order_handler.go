package main

import (
	"database/sql"
	"net/http"
)

type Order struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `phone:"phone"`
}

// GetOrdersHandler is a handler for getting a list of orders
var GetOrdersHandler = AppHandler(func(r *http.Request) *AppRes {
	rows, err := db.Query("select first_name, last_name, email, phone from orders")
	if err != nil {
		return &AppRes{
			Code:  http.StatusInternalServerError,
			Error: newJSONError("Cannot fetch orders", err),
		}
	}
	defer rows.Close()
	var orders []Order
	for rows.Next() {
		o := Order{}
		err := rows.Scan(&o.FirstName, &o.LastName, &o.Email, &o.Phone)
		if err != nil {
			return &AppRes{
				Code:  http.StatusInternalServerError,
				Error: newJSONError("Cannot fetch orders", err),
			}
		}
		orders = append(orders, o)
	}
	return &AppRes{Code: http.StatusOK, Data: orders}
})

type OrderQuery struct {
	Rows []sql.Rows
}
type OrderGetter interface {
	GetOrder(buf Order) error
}

var PostOrderHandler = AppHandler(func(r *http.Request) *AppRes {
	return &AppRes{}
})
