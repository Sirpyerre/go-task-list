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
	mongoDB := os.Getenv("MONGO_DB")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	dns := fmt.Sprintf("mongodb://%s:%v", mongoHost, 27017)
	clientOptions := options.Client().ApplyURI(dns).
		SetAuth(options.Credential{
			AuthSource: mongoDB, Username: mongoUser, Password: mongoPassword,
		})
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
