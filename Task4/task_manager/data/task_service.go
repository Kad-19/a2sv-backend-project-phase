package data

import (
	"fmt"
	"errors"
	"task_manager/models"
	"time"
)

// Mock data for tasks
var tasks = []models.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: models.Pending},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: models.InProgress},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: models.Completed},
}
var tasksCount = 3

// GetTasks returns the list of tasks
func GetTasks() []models.Task {
	return tasks
}

// getTaskByID returns a task by its ID
func GetTaskByID(id string) *models.Task {
	for _, task := range tasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}

// CreateTask adds a new task to the list
func CreateTask(task models.Task) models.Task {
	task.ID = fmt.Sprint(tasksCount + 1)
	task.Status = models.Pending
	tasksCount++
	tasks = append(tasks, task)
	return task
}

// UpdateTask updates an existing task
func UpdateTask(id string, updatedTask models.Task) error{
	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
			return nil
		}
	}

	return errors.New("task not found")
}

// DeleteTask removes a task by its ID
func DeleteTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}