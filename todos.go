package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type TodoStatus struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type User struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type Priority struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type Todo struct {
	ID          string      `json:"id" bson:"_id"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	Status      *TodoStatus `json:"status" bson:"status"`
	User        *User       `json:"user" bson:"user"`
	Priority    *Priority   `json:"priority" bson:"priority"`
	CreatedAt   time.Time   `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time   `json:"updatedAt" bson:"updated_at"`
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	// Get all todos from the database
	coll := client.Database("tododb").Collection("todos")
	ctx := r.Context()
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	todos := []*Todo{}
	for cursor.Next(ctx) {
		t := &Todo{}
		if err := cursor.Decode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	// Send the todos as the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	t := &Todo{}
	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.ID = uuid.New().String()
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	// Insert the todo into the database
	coll := client.Database("tododb").Collection("todos")
	ctx := r.Context()
	_, err = coll.InsertOne(ctx, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the todo as the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the todo from the database
	coll := client.Database("tododb").Collection("todos")
	ctx := r.Context()
	t := &Todo{}
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the todo as the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Parse the request body
	t := &Todo{}
	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the todo in the database
	coll := client.Database("tododb").Collection("todos")
	ctx := r.Context()
	_, err = coll.ReplaceOne(ctx, bson.M{"_id": id}, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the todo as the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the todo from the database
	coll := client.Database("tododb").Collection("todos")
	ctx := r.Context()
	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success message as the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"message": "success"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
