package routes

import (
	"rkoning/angular-example-app/api/controllers"

	"github.com/gorilla/mux"
)

// Routes handles all the routes
func Routes(r *mux.Router) {
	r.HandleFunc("/threads/{id}/comments", controllers.GetAllCommentsForThread).Methods("GET").Schemes("http")
	r.HandleFunc("/threads", controllers.GetAllThreads).Methods("GET").Schemes("http")
	r.HandleFunc("/threads/{id}/comments", controllers.AddCommentHandler).Methods("POST").Schemes("http")
	r.HandleFunc("/threads", controllers.AddThreadHandler).Methods("POST").Schemes("http")
	// r.HandleFunc("/", controllers.PutAPIHandler).Methods("Put").Schemes("http")
	r.HandleFunc("/threads/{id}/comments/{commentId}", controllers.DeleteCommentHandler).Methods("DELETE").Schemes("http")
	r.HandleFunc("/threads/{id}", controllers.DeleteThreadHandler).Methods("DELETE").Schemes("http")
}
