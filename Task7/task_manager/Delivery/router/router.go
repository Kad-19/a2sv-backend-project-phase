package router

import (
	controller "task_manager/Delivery/controllers"
	domain "task_manager/Domain"
	"task_manager/Infrastructure"
	repository "task_manager/Repositories"
	usecase "task_manager/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Setup(env *Infrastructure.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewUserRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	// protectedRouter.Use(Infrastructure.AuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewTaskRouter(env, timeout, db, protectedRouter)
}

func NewTaskRouter(env *Infrastructure.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/tasks", tc.Fetch)
	group.POST("/tasks", tc.Create)
	group.PUT("/tasks", tc.Update)
	group.DELETE("/tasks/:id", tc.Delete)
}

func NewUserRouter(env *Infrastructure.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout, env),
	}
	group.POST("/register", uc.Create)
	group.POST("/login", uc.Login)
}
