package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Thread represents a thread
//
// A Thread is the main component in this application
//
// A thread can have as many comments as the users add
//
// swagger:parameters thread
type Thread struct {
	// the id
	//
	// required: true
	// min: 1
	ID *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// the title
	//
	// required: true
	// min: 1
	Title string `bson:"title,omitempty" json:"title"`
	// the text
	//
	// required: true
	// min: 1
	Text string `bson:"text,omitempty" json:"text"`
	// the id for this user
	//
	// required: true
	// min: 1
	UserID string `bson:"userId,omitempty" json:"userId"`
	// the time in unix the thread was submitted
	//
	// required: true
	// min: 1
	Timestamp int64 `bson:"timestamp,omitempty" json:"timestamp"`
}

// Comment request model
// swagger:parameters comment
type Comment struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string              `bson:"userId,omitempty" json:"userId"`
	ParentID  string              `bson:"parentId,omitempty" json:"parentId"`
	ThreadID  string              `bson:"threadId,omitempty" json:"threadId"`
	Text      string              `bson:"text,omitempty" json:"text"`
	Timestamp int64               `bson:"timestamp,omitempty" json:"timestamp"`
}

// DATABASE INSTANCE
var threadCollection *mongo.Collection

var commentCollection *mongo.Collection

// ThreadsCollection is a collection in mongo
func ThreadsCollection(c *mongo.Database) {
	threadCollection = c.Collection("threads")
}

// CommentsCollection is a collection of comments
func CommentsCollection(c *mongo.Database) {
	commentCollection = c.Collection("comments")
}

// GetAllThreads Get Handler
func GetAllThreads(w http.ResponseWriter, r *http.Request) {
	threads := []Thread{}

	cursor, err := threadCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var thread Thread
		cursor.Decode(&thread)
		threads = append(threads, thread)
	}

	json := simplejson.New()
	json.Set("data", threads)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

// GetThread gets thread
func GetThread(w http.ResponseWriter, r *http.Request) {
	var thread Thread

	cursor := threadCollection.FindOne(context.TODO(), bson.M{})

	// Iterate through the returned cursor.

	cursor.Decode(&thread)

	json := simplejson.New()
	json.Set("data", thread)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

// GetAllCommentsForThread Get Handler
func GetAllCommentsForThread(w http.ResponseWriter, r *http.Request) {
	threadID := mux.Vars(r)["id"]
	comments := []Comment{}

	cursor, err := commentCollection.Find(context.TODO(), bson.M{"$or": []interface{}{
		bson.M{"threadId": threadID},
		bson.M{"parentId": threadID}}})
	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var comment Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}

	json := simplejson.New()
	json.Set("data", comments)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

// AddThreadHandler for post functions
func AddThreadHandler(w http.ResponseWriter, r *http.Request) {
	var t Thread
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.Timestamp = time.Now().Unix()

	result, insertErr := threadCollection.InsertOne(context.TODO(), t)

	if insertErr != nil {
		log.Println(insertErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Payload(insertErr))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(Payload(result))
	return
}

// AddCommentHandler for post functions
func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	threadID := mux.Vars(r)["id"]
	var c Comment
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.ThreadID = threadID
	c.Timestamp = time.Now().Unix()

	result, insertErr := commentCollection.InsertOne(context.TODO(), c)

	if insertErr != nil {
		log.Println(insertErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Payload(insertErr))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(Payload(result))
	return
}

// // PutAPIHandler for put functions
// func PutAPIHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Api Retrieved")
// }

// DeleteThreadHandler for delete functions
func DeleteThreadHandler(w http.ResponseWriter, r *http.Request) {
	threadID := mux.Vars(r)["id"]
	docID, err := primitive.ObjectIDFromHex(threadID)
	json := simplejson.New()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Payload(err))
		return
	}

	delComments, delCommentsErr := commentCollection.DeleteMany(context.TODO(), bson.M{"$or": []interface{}{
		bson.M{"threadId": threadID},
		bson.M{"parentId": threadID}}})

	if delCommentsErr != nil {
		log.Println(delCommentsErr)
		return
	}

	json.Set("commentsDeleted", delComments)

	result, insertErr := threadCollection.DeleteOne(context.TODO(), bson.M{"_id": docID})

	if insertErr != nil {
		log.Println(insertErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Payload(insertErr))
		return
	}

	json.Set("threadDeleted", result)

	w.WriteHeader(http.StatusOK)
	w.Write(Payload(json))
	return
}

// DeleteCommentHandler for delete functions
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	commentID := mux.Vars(r)["commentId"]
	docID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Payload(err))
		return
	}
	result, insertErr := commentCollection.DeleteOne(context.TODO(), bson.M{"_id": docID})

	if insertErr != nil {
		log.Println(insertErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Payload(insertErr))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(Payload(result))
	return
}

// Payload generates payload for request response
func Payload(data interface{}) (result []byte) {
	json := simplejson.New()
	json.Set("data", data)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	return payload
}

// func GetAllTodos(c *gin.Context) {
// 	todos := []Todo{}
// 	cursor, err := collection.Find(context.TODO(), bson.M{})

// 	if err != nil {
// 		log.Printf("Error while getting all todos, Reason: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Something went wrong",
// 		})
// 		return
// 	}

// 	// Iterate through the returned cursor.
// 	for cursor.Next(context.TODO()) {
// 		var todo Todo
// 		cursor.Decode(&todo)
// 		todos = append(todos, todo)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "All Todos",
// 		"data":    todos,
// 	})
// 	return
// }

// func CreateTodo(c *gin.Context) {
// 	var todo Todo
// 	c.BindJSON(&todo)
// 	title := todo.Title
// 	body := todo.Body
// 	completed := todo.Completed
// 	id := guuid.New().String()

// 	newTodo := Todo{
// 		ID:        id,
// 		Title:     title,
// 		Body:      body,
// 		Completed: completed,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	_, err := collection.InsertOne(context.TODO(), newTodo)

// 	if err != nil {
// 		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Something went wrong",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"status":  http.StatusCreated,
// 		"message": "Todo created Successfully",
// 	})
// 	return
// }

// func GetSingleTodo(c *gin.Context) {
// 	todoId := c.Param("todoId")

// 	todo := Todo{}
// 	err := collection.FindOne(context.TODO(), bson.M{"id": todoId}).Decode(&todo)
// 	if err != nil {
// 		log.Printf("Error while getting a single todo, Reason: %v\n", err)
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  http.StatusNotFound,
// 			"message": "Todo not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Single Todo",
// 		"data":    todo,
// 	})
// 	return
// }

// func EditTodo(c *gin.Context) {
// 	todoId := c.Param("todoId")
// 	var todo Todo
// 	c.BindJSON(&todo)
// 	completed := todo.Completed

// 	newData := bson.M{
// 		"$set": bson.M{
// 			"completed":  completed,
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": todoId}, newData)
// 	if err != nil {
// 		log.Printf("Error, Reason: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  500,
// 			"message": "Something went wrong",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"message": "Todo Edited Successfully",
// 	})
// 	return
// }

// func DeleteTodo(c *gin.Context) {
// 	todoId := c.Param("todoId")

// 	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": todoId})
// 	if err != nil {
// 		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Something went wrong",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Todo deleted successfully",
// 	})
// }
