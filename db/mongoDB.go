package db

import (
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBconnect -> connects mongo
func ConnectDB() *mongo.Database {

	envFile, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Failed to load env")
	}
	clientOptions := options.Client().ApplyURI(envFile["MONGO_URL"])
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Connection Failed to MongoDB")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Ping Failed to MongoDB")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	db := client.Database(envFile["MONGO_DB_NAME"])
	return db
}
