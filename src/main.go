package main

import (
	"crmbackend/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", controller.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controller.GetCustomer).Methods("GET")

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", router)
}
