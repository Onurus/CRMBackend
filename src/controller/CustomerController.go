package controller

import (
	"crmbackend/model"
	"crmbackend/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	if customer.ID == headerId {
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(customer)
	} else {
		responseWriter.WriteHeader(http.StatusNotFound)
	}
}

func CreateCustomer(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Print(err)
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	var customer model.Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		fmt.Print(err)
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	customer = service.CreateCustomer(customer)
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(customer)
}

func DeleteCustomer(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	headerId := getIdFromRequestHeader(request)
	customer := service.GetAllCustomer(headerId)

	if customer.ID == headerId {
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
