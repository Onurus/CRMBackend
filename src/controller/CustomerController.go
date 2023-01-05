package controller

import (
	"crmbackend/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCustomers(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(service.GetAllCustomers())
}

func GetCustomer(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	item := mux.Vars(request)
	headerId, _ := strconv.ParseInt(item["id"], 10, 64)
	customer := service.GetAllCustomer(headerId)

	if customer.Id == headerId {
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(customer)
	} else {
		responseWriter.WriteHeader(http.StatusNotFound)
	}
}

func CreateCustomer(responseWriter http.ResponseWriter, request *http.Request) {

}

func DeleteCustomer(responseWriter http.ResponseWriter, request *http.Request) {

}

func UpdateCustomer(responseWriter http.ResponseWriter, request *http.Request) {

}
