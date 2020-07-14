package routes

import (
	"rkoning/angular-example-app/api/controllers"

	"github.com/gorilla/mux"
)

// Routes handles all the routes
func Routes(r *mux.Router) {
	r.HandleFunc("/", controllers.GetAPIHandler).Methods("GET").Schemes("http")
	r.HandleFunc("/", controllers.PostAPIHandler).Methods("POST").Schemes("http")
	r.HandleFunc("/", controllers.PutAPIHandler).Methods("Put").Schemes("http")
	r.HandleFunc("/", controllers.DeleteAPIHandler).Methods("DELETE").Schemes("http")
}

// func welcome(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"message": "Welcome To API",
// 	})
// 	return
// }

// func notFound(c *gin.Context) {
// 	c.JSON(http.StatusNotFound, gin.H{
// 		"status":  404,
// 		"message": "Route Not Found",
// 	})
// 	return
// }
