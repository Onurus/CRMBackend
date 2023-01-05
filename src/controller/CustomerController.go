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
	headerId := getIdFromRequestHeader(request)
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
	responseWriter.Header().Set("Content-Type", "application/json")
	headerId := getIdFromRequestHeader(request)
	customer := service.GetAllCustomer(headerId)

	if customer.Id == headerId {
		service.DeleteCustomer(headerId)
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(service.GetAllCustomers())
	} else {
		responseWriter.WriteHeader(http.StatusNotFound)
	}
}

func UpdateCustomer(responseWriter http.ResponseWriter, request *http.Request) {

}

func getIdFromRequestHeader(request *http.Request) int64 {
	item := mux.Vars(request)
	headerId, _ := strconv.ParseInt(item["id"], 10, 64)
	return headerId
}
