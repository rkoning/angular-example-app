// Basic Message Board API.
//
//
// Terms Of Service:
//
//     Schemes: http
//     Host: localhost:10000
//     Version: 1.0.0
//     Contact: Joel Gilbert<youwish@email.com>
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

	"github.com/gorilla/handlers"

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

	fs := http.FileServer(http.Dir("./swaggerui"))
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", fs))
	r.PathPrefix("/swagger").Handler(http.StripPrefix("/swagger/", fs))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
