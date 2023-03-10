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

// Expected API methods
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
		customerNotFound(responseWriter)
	}
}

func AddCustomer(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var customer model.Customer
	if validateRequestAndParseBody(responseWriter, request, &customer) {
		customer = service.CreateCustomer(customer)
		responseWriter.WriteHeader(http.StatusCreated)
		json.NewEncoder(responseWriter).Encode(customer)
	}
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
		customerNotFound(responseWriter)
	}
}

func UpdateCustomer(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	headerId := getIdFromRequestHeader(request)
	existingCustomer := service.GetAllCustomer(headerId)

	if existingCustomer.ID == headerId {
		var incomingCustomer model.Customer
		if validateRequestAndParseBody(responseWriter, request, &incomingCustomer) {
			incomingCustomer = service.UpdateCustomer(incomingCustomer, headerId)
			responseWriter.WriteHeader(http.StatusCreated)
			json.NewEncoder(responseWriter).Encode(incomingCustomer)
		}
	} else {
		customerNotFound(responseWriter)
	}

}

// Helper functions
func customerNotFound(responseWriter http.ResponseWriter) {
	responseWriter.WriteHeader(http.StatusNotFound)
	json.NewEncoder(responseWriter).Encode(map[string]string{"message": "Customer Not Found!"})
}

func getIdFromRequestHeader(request *http.Request) int64 {
	item := mux.Vars(request)
	headerId, _ := strconv.ParseInt(item["id"], 10, 64)
	return headerId
}

func validateRequestAndParseBody(responseWriter http.ResponseWriter, request *http.Request, customerRef *model.Customer) bool {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Print(err)
		responseWriter.WriteHeader(http.StatusBadRequest)
		return false
	}
	err = json.Unmarshal(body, customerRef)
	if err != nil {
		fmt.Print(err)
		responseWriter.WriteHeader(http.StatusBadRequest)
		return false
	}
	return true
}
