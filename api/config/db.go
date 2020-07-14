package config

import (
	"context"
	"log"
	"rkoning/angular-example-app/api/controllers"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect is used to connect to the database
func Connect() {
	// Database Config ewy44aEVAiOk2kd3
	clientOptions := options.Client().ApplyURI("mongodb+srv://user:ewy44aEVAiOk2kd3@cluster0.5hlas.gcp.mongodb.net/demo?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("demo")
	controllers.ThreadsCollection(db)
	controllers.CommentsCollection(db)
	return
}
