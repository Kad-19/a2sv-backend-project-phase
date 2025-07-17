package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Status string

const (
	Pending    Status = "Pending"
	InProgress Status = "In Progress"
	Completed  Status = "Completed"
)

type Task struct {
 ID          bson.ObjectID    `json:"id" bson:"_id,omitempty"`
 Title       string    `json:"title" binding:"required"`
 Description string    `json:"description" binding:"required"`
 DueDate     time.Time `json:"dueDate" binding:"required"`
 Status      Status    `json:"status" enum:"Pending,In Progress,Completed" `
}

