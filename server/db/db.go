package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func DBInstance() *mongo.Client {
	fmt.Println("Connecting to DB")
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	monogo_uri := os.Getenv("MONGODB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(monogo_uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
}
