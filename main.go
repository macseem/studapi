package main

import (
	"github.com/macseem/studapi/connectors/places"
	"github.com/macseem/studapi/connectors/shipping"
	//"github.com/macseem/studapi/studentcom"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var config = &Config{}
var httpClient *http.Client = &http.Client{}
var db *sql.DB
var placesClient *places.Client
var shippingClient *shipping.Client
func init() {
	f, err := os.Open("./config.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(config)
	if err != nil {
		panic(err)
	}

	placesClient = &places.Client{
		APIKey:     config.GoogleAPIKey,
		HTTPGetter: httpClient,
	}
	shippingClient = &shipping.Client{
		APIKey:     config.ShippingForceAPIKey,
		HTTPPoster: httpClient,
	}

	db, err = sql.Open("postgres", config.DBConnStr)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	// http.HandleFunc("/getlistacc", HandlerGetListAcc)
	// http.HandleFunc("/bookaccomodation", HandlerPostBook)
	http.Handle("/places", PlacesHandler)
	http.Handle("/shipping/prices", ShippingPricesHandler)
	http.Handle("/orders", getOrdersHandler)
	http.Handle("/shipping/order", shippingOrderHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// func HandlerGetListAcc(w http.ResponseWriter, r *http.Request) {
// 	var reqProp studentcom.GetListAccProperties
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&reqProp)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"Error": "TODO add message here"}`))
// 		return
// 	}
// 	list, err := GetListAcc(reqProp)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"Error": "TODO add message here"}`))
// 		return
// 	}
// 	respBody, err := json.Marshal(list)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"Error": "TODO add message here"}`))
// 		return
// 	}
// 	w.WriteHeader(200)
// 	w.Write(respBody)
// }
//
// func GetListAcc(reqProp studentcom.GetListAccProperties) (studentcom.PropertiesResponse, error) {
// 	result, err := studentcom.GetListAcc(reqProp)
// 	if err != nil {
// 		log.Println("Error: " + err.Error())
// 		return result, err
// 	}
// 	return result, err
// }
//
// type Invoice struct{}
//
// //TODO
// func HandlerPostBook(w http.ResponseWriter, r *http.Request) {
// 	var newLead studentcom.NewLead
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&newLead)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"Error": "TODO add message here"}`))
// 		return
// 	}
// 	invoice, err := BookToInvoice(newLead)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"Error": "TODO add message here"}`))
// 		return
// 	}
// 	respBody, err := json.Marshal(invoice)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"Error": "TODO add message here"}`))
// 		return
// 	}
// 	w.WriteHeader(200)
// 	w.Write(respBody)
// }
//
// func BookToInvoice(lead studentcom.NewLead) (invoice Invoice, err error) {
// 	resp, err := studentcom.PostBook(lead)
// 	if err != nil {
// 		return
// 	}
// 	invoice, err = ConvertRespToInvoice(resp)
// 	if err != nil {
// 		return
// 	}
// 	err = writeLeadToDB(lead, resp)
// 	if err != nil {
// 		return
// 	}
// 	return
// }
//
// //TODO
// func ConvertRespToInvoice(data studentcom.LeadResponse) (invoice Invoice, err error) {
// 	return
// }
//
// //TODO
// func writeLeadToDB(lead studentcom.NewLead, resp studentcom.LeadResponse) error {
// 	return nil
// }
