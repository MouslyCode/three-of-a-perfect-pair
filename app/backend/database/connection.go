package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var TaskCollection *mongo.Collection

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Load .env !")
	}
	uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	db := client.Database("task-db")

	TaskCollection = db.Collection("tasks")

	fmt.Println("MongoDB Connect Succesfully!")
}
