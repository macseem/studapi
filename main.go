package main

import (
	"clearmove/studapi/studentcom"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getlistacc", HandlerGetListAcc)
	http.HandleFunc("/bookaccomodation", HandlerPostBook)
	// TODO: Move config values to .env

	http.Handle(
		"/places",
		NewPlacesHandler("AIzaSyBxbE8g9mKeNXvgERMb_moW2weunJNu1X4"),
	)

	http.Handle(
		"/shipping/prices",
		NewShippingPricesHandler(
			"a0d520de-957a-4e54-af58-b936647f0387",
			"AIzaSyBxbE8g9mKeNXvgERMb_moW2weunJNu1X4",
		),
	)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func HandlerGetListAcc(w http.ResponseWriter, r *http.Request) {
	var reqProp studentcom.GetListAccProperties
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqProp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error": "TODO add message here"}`))
		return
	}
	list, err := GetListAcc(reqProp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error": "TODO add message here"}`))
		return
	}
	respBody, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error": "TODO add message here"}`))
		return
	}
	w.WriteHeader(200)
	w.Write(respBody)
}

func GetListAcc(reqProp studentcom.GetListAccProperties) (studentcom.PropertiesResponse, error) {
	result, err := studentcom.GetListAcc(reqProp)
	if err != nil {
		log.Println("Error: " + err.Error())
		return result, err
	}
	return result, err
}

type Invoice struct{}

//TODO
func HandlerPostBook(w http.ResponseWriter, r *http.Request) {
	var newLead studentcom.NewLead
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newLead)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error": "TODO add message here"}`))
		return
	}
	invoice, err := BookToInvoice(newLead)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error": "TODO add message here"}`))
		return
	}
	respBody, err := json.Marshal(invoice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Error": "TODO add message here"}`))
		return
	}
	w.WriteHeader(200)
	w.Write(respBody)
}

func BookToInvoice(lead studentcom.NewLead) (invoice Invoice, err error) {
	resp, err := studentcom.PostBook(lead)
	if err != nil {
		return
	}
	invoice, err = ConvertRespToInvoice(resp)
	if err != nil {
		return
	}
	err = writeLeadToDB(lead, resp)
	if err != nil {
		return
	}
	return
}

//TODO
func ConvertRespToInvoice(data studentcom.LeadResponse) (invoice Invoice, err error) {
	return
}

//TODO
func writeLeadToDB(lead studentcom.NewLead, resp studentcom.LeadResponse) error {
	return nil
}
