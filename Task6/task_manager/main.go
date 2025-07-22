package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
  
	data.InitMongoDB()

	// Initialize the router
	r := router.SetupRouter()

	// Start the server on port 8080
	if err := r.Run("localhost:8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}

