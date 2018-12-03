package main

import "errors"

type Customer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GenerateCustomers() []Customer {
	customers := []Customer{
		Customer{Id: 1, Name: "Mahmut"},
		Customer{Id: 2, Name: "Tuncer"},
	}
	return customers
}

func GetSingleCustomer(customerId int) (*Customer, error) {
	if customerId > 2 {
		return nil, errors.New("Argument error")
	}
	customers := GenerateCustomers()
	customer := customers[customerId-1]
	return &customer, nil
}
