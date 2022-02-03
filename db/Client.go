package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() (client *mongo.Client) {
	mongoHost := os.Getenv("MONGO_HOST")

	// Without authentication
	dns := fmt.Sprintf("mongodb://%s:%v", mongoHost, 27017)
	clientOptions := options.Client().ApplyURI(dns)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Client Error:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Ping error: ", err)
	}

	return
}
