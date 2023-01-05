package service

import (
	"crmbackend/model"

	"golang.org/x/exp/maps"
)

var customerMap = map[int64]model.Customer{
	1: {
		Id:        1,
		Name:      "Onur Usta",
		Role:      "Software Engineer",
		Email:     "onurusta@gmail.com",
		Phone:     "3661278",
		Contacted: true,
	},
	2: {
		Id:        2,
		Name:      "Erkut Yuzgec",
		Role:      "Software Engineer",
		Email:     "erkutYz@gmail.com",
		Phone:     "3661161",
		Contacted: true,
	},
	3: {
		Id:        3,
		Name:      "Big Boss",
		Role:      "Principal Engineer",
		Email:     "bb@gmail.com",
		Phone:     "4414596",
		Contacted: true,
	},
	4: {
		Id:        4,
		Name:      "Tugce Usta",
		Role:      "DevOps Engineer",
		Email:     "tugceUsta@gmail.com",
		Phone:     "5327764108",
		Contacted: false,
	},
	5: {
		Id:        5,
		Name:      "Doga Usta",
		Role:      "Junior Software Engineer",
		Email:     "dogaUsta@gmail.com",
		Phone:     "NotAvaible",
		Contacted: false,
	},
}

func GetAllCustomers() []model.Customer {
	return maps.Values(customerMap)
}

func GetAllCustomer(id int64) model.Customer {
	return customerMap[id]
}
