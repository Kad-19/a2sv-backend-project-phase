package data

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

var tasks_collection *mongo.Collection
var user_collection *mongo.Collection

func InitMongoDB() {
	// Load environment variables from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found or error loading it")
	}
	mongoAPIURL := os.Getenv("MONGO_API_URL")
	
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoAPIURL).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 	panic(err)
	// 	}
	// }()

	// Initialize the tasks collection
	tasks_collection = client.Database("task_manager").Collection("tasks")

	// Initialize the users collection
	user_collection = client.Database("task_manager").Collection("users")

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
