package main

// Config a struct of configuration
type Config struct {
  GoogleAPIKey string `json:"google_api_key"`
  ShippingForceAPIKey string `json:"shipping_force_api_key"`
  DBConnStr string `json:"db_conn_str"`
}
