package main
import (
  "net/http"
)
type Order struct {}
// GetOrdersHandler is a handler for getting a list of orders
var GetOrdersHandler = AppHandler(func(r *http.Request) (interface{}, *jsonErrorRes) {
  return Order{}, nil
})

var PostOrderHandler = AppHandler(func(r *http.Request) (interface{}, *jsonErrorRes) {
  return Order{}, nil
})
