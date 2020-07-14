package main

import (
	"log"
	"net/http"

	"rkoning/angular-example-app/api/config"
	"rkoning/angular-example-app/api/routes"

	"github.com/gorilla/mux"
)

func main() {

	// Database
	config.Connect()

	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	routes.Routes(r)

	log.Fatal(http.ListenAndServe(":10000", r))
}
