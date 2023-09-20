package main

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        string    `bson:"_id"`
	Task      string    `bson:"task"`
	Completed bool      `bson:"completed"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type TodoDao interface {
	GetAll() ([]*Todo, error)
	Get(id string) (*Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id string) error
}

type TodoDaoImpl struct {
	client *mongo.Client
}

func NewTodoDao(client *mongo.Client) TodoDao {
	return &TodoDaoImpl{client}
}

func (dao *TodoDaoImpl) GetAll() ([]*Todo, error) {
	ctx := context.Background()
	// get all todos from mongodb
	coll := dao.client.Database("todo").Collection("todos")
	cur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	// convert to []*Todo
	var todos []*Todo
	for cur.Next(ctx) {
		var todo *Todo
		if err := cur.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (dao *TodoDaoImpl) Get(id string) (*Todo, error) {
	todo := &Todo{}
	coll := dao.client.Database("todo").Collection("todos")
	err := coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (dao *TodoDaoImpl) Create(todo *Todo) error {
	todo.ID = uuid.New().String()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	_, err := dao.client.Database("todo").Collection("todos").InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDaoImpl) Update(todo *Todo) error {
	return errors.New("not implemented yet")
}

func (dao *TodoDaoImpl) Delete(id string) error {
	return errors.New("not implemented yet")
}
