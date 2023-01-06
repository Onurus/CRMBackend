# CRMBackend
CRM Backend Go Project for Udacity

go version 1.19.3

# Used Livraries

- github.com/gorilla/mux
    Mux library is used for routing and registering the handler functions to endpoints.

- golang.org/x/exp
    exp library is just used for map functionality. getting all values of a map is much easier with exp libary.

# How to run
go-to src folder
run "go run main.go" 

# Usage

After running the main.go the server start on localhost at port 3000.
You can try the following API's (listed below) through Postman or any other tool

### Services List

- Getting all customers through a the /customers path 
    "/customers" ("GET")

- Getting a single customer through a /customers/{id} path
    "/customers/{id}" ("GET")

-  Creating a customer through a /customers path 
    "/customers" ("POST")

-  Updating a customer through a /customers/{id} path 
    "/customers/{id}" ("DELETE")

- Deleting a customer through a /customers/{id} path
    "/customers/{id}" ("PATCH")

### Code structure
- Mock DB ( map impl) -> service/MockCustomerDB.go ( 5 records at initial )
- Customer model -> model/Customer.go
- Request handlers -> controller/CustomerController.go
- API registeration, Port defs. -> main.go

