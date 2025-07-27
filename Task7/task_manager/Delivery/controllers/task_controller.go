package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task_manager/Domain"
)

type TaskController struct {
	TaskUsecase Domain.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	var task Domain.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	err = tc.TaskUsecase.Create(c, &task, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("user_id")

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) Update(c *gin.Context) {
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := tc.TaskUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := tc.TaskUsecase.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
