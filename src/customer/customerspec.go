package main

// GetCustomerQueryParams model.
//
// swagger:parameters getCustomer
type GetCustomerQueryParams struct {
	// CustomerId to query.
	//
	// required: true
	// minimum: 1
	CustomerId int32 `json:"customerId"`
}

// CustomerResponse response model
//
// This is used for returning a response with a single customer as body
//
// swagger:response customerResponse
type CustomerResponse struct {
	// out: body
	Payload *Customer `json:"customer"`
}
