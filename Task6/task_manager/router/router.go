package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the router and defines the routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	protected := router.Group("", middleware.AuthMiddleware())

	// Task routes
	protected.GET("/tasks", controllers.GetTasks)
	protected.GET("/tasks/:id", controllers.GetTaskByID)
	protected.POST("/tasks", middleware.OnlyAdmin(), controllers.CreateTask)
	protected.PUT("/tasks/:id", middleware.OnlyAdmin(), controllers.UpdateTask)
	protected.DELETE("/tasks/:id", middleware.OnlyAdmin(), controllers.DeleteTask)

	// User routes	
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)


	return router
}