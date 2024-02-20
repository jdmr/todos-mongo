package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Create a new context
	ctx := context.Background()
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	// setting up database collections
	priorities := []Priority{
		{ID: "1", Name: "Low"},
		{ID: "2", Name: "Medium"},
		{ID: "3", Name: "High"},
	}
	// upsert priorities
	coll := client.Database("tododb").Collection("priorities")
	opts := options.Update().SetUpsert(true)
	for _, p := range priorities {
		filter := bson.M{"_id": p.ID}
		update := bson.M{"$set": p}
		_, err := coll.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	statuses := []TodoStatus{
		{ID: "1", Name: "Not Started"},
		{ID: "2", Name: "In Progress"},
		{ID: "3", Name: "Completed"},
	}
	// upsert statuses
	coll = client.Database("tododb").Collection("statuses")
	for _, s := range statuses {
		filter := bson.M{"_id": s.ID}
		update := bson.M{"$set": s}
		_, err := coll.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/todos", getAllTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	// Start the server
	log.Println("Started the server on :8080...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("server stopped with error: ", err)
	}
}
