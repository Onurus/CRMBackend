package controller

import (
	"crmbackend/service"
	"encoding/json"
	"net/http"
)

func GetCustomers(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(service.GetAllCustomers())
}
