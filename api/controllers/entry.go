package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Thread is a main thread
type Thread struct {
	ID     string `bson:"_id" json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

// Comment is a reply to the Thread
type Comment struct {
	ID       string `bson:"_id" json:"id"`
	UserID   string `json:"userId"`
	ParentID string `json:"parentId"`
	Text     string `json:"text"`
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

	_, err := threadCollection.InsertOne(context.TODO(), bson.D{
		{Key: "Title", Value: "Inserted From controller?"},
		{Key: "Text", Value: "Wow, this was inserted from the controller!"},
		{Key: "UserID", Value: "1234"}})

	if err != nil {
		log.Printf("Error Saving Data, Reason %v\n", err)
	}

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

// PostAPIHandler for post functions
func PostAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Api Retrieved")
}

// PutAPIHandler for put functions
func PutAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Api Retrieved")
}

// DeleteAPIHandler for delete functions
func DeleteAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Api Retrieved")
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
