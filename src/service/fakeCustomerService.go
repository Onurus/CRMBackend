package service

import "crmbackend/model"

var customers = []model.Customer{
	{
		Id:        1,
		Name:      "Onur Usta",
		Role:      "Software Engineer",
		Email:     "onurusta@gmail.com",
		Phone:     "3661278",
		Contacted: true,
	},
	{
		Id:        2,
		Name:      "Erkut Yuzgec",
		Role:      "Software Engineer",
		Email:     "erkutYz@gmail.com",
		Phone:     "3661161",
		Contacted: true,
	},
	{
		Id:        3,
		Name:      "Big Boss",
		Role:      "Principal Engineer",
		Email:     "bb@gmail.com",
		Phone:     "4414596",
		Contacted: true,
	},
	{
		Id:        4,
		Name:      "Tugce Usta",
		Role:      "DevOps Engineer",
		Email:     "tugceUsta@gmail.com",
		Phone:     "5327764108",
		Contacted: false,
	},
	{
		Id:        5,
		Name:      "Doga Usta",
		Role:      "Junior Software Engineer",
		Email:     "dogaUsta@gmail.com",
		Phone:     "NotAvaible",
		Contacted: false,
	},
}

func GetAllCustomers() []model.Customer {
	return customers
}
