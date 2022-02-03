package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (task *Task) CreateTask() (insertResult *mongo.InsertOneResult) {
	client := Connection()
	collection := client.Database("TasksList").Collection("tasks")
	defer client.Disconnect(context.TODO())

	insertResult, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Println("Error inserting", err)
	}

	return
}

func (task *Task) GetTasks(limit int64) (results []*Task) {
	client := Connection()
	collection := client.Database("TasksList").Collection("tasks")
	defer client.Disconnect(context.TODO())

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Println("Error finding ", err)
	}

	for cur.Next(context.TODO()) {
		var item Task
		err2 := cur.Decode(item)
		if err2 != nil {
			log.Println("Error decoding ", err2)
		}

		results = append(results, &item)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return
}
