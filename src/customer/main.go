// Customer service rewritten in Go.
//
// This service can be used to interact customer domain of the Westeros Shop
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	fs := http.FileServer(http.Dir("./swagger/"))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))

	log.Fatal(http.ListenAndServe(":6543", router))
}
