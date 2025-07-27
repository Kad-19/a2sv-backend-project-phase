package controller

import (
	"net/http"
	"task_manager/Domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase Domain.UserUsecase
}
type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type CreateUserDTO struct {
	ID       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

func (uc *UserController) Create(c *gin.Context) {
	var dto CreateUserDTO

	err := c.ShouldBind(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.UserUsecase.Create(c, uc.ChangeToDomain(&dto))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var request LoginDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := uc.UserUsecase.Login(c, request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user, "access_token": token})
}

func (uc *UserController) ChangeToDomain(userDTO *CreateUserDTO) *Domain.User {
	var user Domain.User
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.Role = userDTO.Role
	return &user
}