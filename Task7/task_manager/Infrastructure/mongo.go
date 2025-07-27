package Infrastructure

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
  	"go.mongodb.org/mongo-driver/v2/mongo/options"
  	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func InitMongoDB(mongoAPIURL string) *mongo.Database {

	// check if mongoAPIURL is set
	if mongoAPIURL == "" {
		panic("MONGO_API_URL environment variable is not set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoAPIURL).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// Initialize the Database variable
	Database := client.Database("task_manager")

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return Database
}
