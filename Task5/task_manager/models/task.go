package models

import "time"

type Status string

const (
	Pending    Status = "Pending"
	InProgress Status = "In Progress"
	Completed  Status = "Completed"
)

type Task struct {
 ID          string    `json:"id"`
 Title       string    `json:"title" binding:"required"`
 Description string    `json:"description" binding:"required"`
 DueDate     time.Time `json:"dueDate" binding:"required"`
 Status      Status    `json:"status" enum:"Pending,In Progress,Completed" `
}

