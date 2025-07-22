package controllers

import (
	"fmt"
	"net/http"
	"os"
	"task_manager/data"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve tasks: %v", err)})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task := data.GetTaskByID(id)
	if task == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid task data: %v", err)})
		return
	} else {
		newTask, err := data.CreateTask(newTask)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create task: %v", err)})
			return
		}
		c.IndentedJSON(http.StatusCreated, newTask)
	}
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}
	if err := data.UpdateTask(id, updatedTask); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, updatedTask)
	}
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := data.DeleteTask(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusNoContent, nil)
	}
}

func RegisterUser(c *gin.Context) {
  var user models.User
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload" + err.Error()})
    return
  }
  // Check if user already exists
  existingUser, err := data.GetUserByEmail(user.Email)
  if err == nil && existingUser != nil {
    c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
    return
  }

  // User registration logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user.Password = string(hashedPassword)
	created_user, err := data.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": created_user})
}

func LoginUser(c *gin.Context) {
  var user models.LoginPayload
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
    return
  }

  // TODO: Implement user login logic

  // User login logic
	existingUser, err := data.GetUserByEmail(user.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": existingUser.ID,
	"email":   existingUser.Email,
	"role":    existingUser.Role,
	"exp":     jwt.TimeFunc().Add(time.Hour * 24).Unix(), // Token expiration time
	"iat":     jwt.TimeFunc().Unix(), // Issued at time
	})

	// Load jwt secret from environment variable
	_ = godotenv.Load()
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set in .env file")
	}
	jwtSecret := []byte(secret)

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": jwtToken})

}