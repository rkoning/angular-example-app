package routes

import (
	"rkoning/angular-example-app/api/controllers"

	"github.com/gorilla/mux"
)

// Routes handles all the routes
func Routes(r *mux.Router) {
	// swagger:operation GET /threads/{id}/comments comments getComments
	// ---
	// summary: Return comments for thread.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the thread
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "schema":
	//       "data":
	//         "type": object
	r.HandleFunc("/threads/{id}/comments", controllers.GetAllCommentsForThread).Methods("GET").Schemes("http")
	// swagger:operation GET /threads threads getThreads
	// ---
	// summary: Return thread.
	// responses:
	//   "200":
	//     "schema":
	//       "data":
	//         "type": object
	r.HandleFunc("/threads", controllers.GetAllThreads).Methods("GET").Schemes("http")
	// swagger:operation POST /threads/{id}/comments comments createComment
	// ---
	// summary: Create Comment
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the thread
	//   type: string
	//   required: true
	// - name: comment
	//   description: comment to add to the thread
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/comment"
	// responses:
	//   "200":
	//     "schema":
	//       "data":
	//         "type": object
	r.HandleFunc("/threads/{id}/comments", controllers.AddCommentHandler).Methods("POST").Schemes("http")
	// swagger:operation POST /threads threads createThread
	// ---
	// summary: Return thread.
	// parameters:
	// - name: comment
	//   description: comment to add to the thread
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/thread"
	// responses:
	//   "201":
	//     "schema":
	//       "data":
	//         "type": object
	r.HandleFunc("/threads", controllers.AddThreadHandler).Methods("POST").Schemes("http")
	// swagger:operation DELETE /threads/{id}/comments comments deleteComment
	// ---
	// summary: Return comments for thread.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the thread
	//   type: string
	//   required: true
	// responses:
	//   "201":
	//     "schema":
	//       "data":
	//         "type": object
	r.HandleFunc("/threads/{id}/comments/{commentId}", controllers.DeleteCommentHandler).Methods("DELETE").Schemes("http")
	// swagger:operation DELETE /threads/{id} threads deleteThread
	// ---
	// summary: Return deleted result
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the thread
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "schema":
	//       "data":
	//         "type": object
	r.HandleFunc("/threads/{id}", controllers.DeleteThreadHandler).Methods("DELETE").Schemes("http")
}
