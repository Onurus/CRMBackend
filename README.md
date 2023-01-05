# CRMBackend
CRM Backend Go Project for Udacity

go version 1.19.3
github.com/gorilla/mux
golang.org/x/exp

# How to run
go-to src folder
run "go run main.go" 

# Usage

After running the main.go the server start on localhost at port 3000.
You can try the following API's (listed below) through Postman or any other tool

-  "/customers" ("GET")
-  "/customers/{id}" ("GET")
-  "/customers" ("POST")
-  "/customers/{id}" ("DELETE")
-  "/customers/{id}" ("PATCH")

# Code structure

- Mock DB ( map impl) -> service/MockCustomerDB.go ( 5 records at initial )
- Customer model -> model/Customer.go
- Request handlers -> controller/CustomerController.go
- API registeration, Port defs. -> main.go

