package main

import (
	"task_manager/Delivery/router"
	"task_manager/Infrastructure"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load environment variables
	env := Infrastructure.NewEnv()
	// Initialize MongoDB connection
	db := Infrastructure.InitMongoDB(env.MongoAPIURL)


	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	// Initialize the router
	router.Setup(env, timeout, *db, gin)

	// Start the server on port 8080
	if err := gin.Run("localhost:8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
	
}

