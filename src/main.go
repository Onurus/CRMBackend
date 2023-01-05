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
	router.HandleFunc("/customers", controller.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controller.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", controller.UpdateCustomer).Methods("PATCH")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../static"))))

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", router)
}
