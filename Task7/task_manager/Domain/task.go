package Domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionTask = "tasks"
)

type Status string

const (
	Pending    Status = "Pending"
	InProgress Status = "In Progress"
	Completed  Status = "Completed"
)

type Task struct {
 ID          bson.ObjectID    `json:"id" bson:"_id,omitempty"`
 UserID 	 bson.ObjectID    `json:"userId" bson:"user_id,omitempty"`
 Title       string    `json:"title" binding:"required"`
 Description string    `json:"description" binding:"required"`
 DueDate     time.Time `json:"dueDate" binding:"required"`
 Status      Status    `json:"status" enum:"Pending,In Progress,Completed" `
}

type TaskRepository interface {
	Create(c context.Context, task *Task, userID string) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
	Update(c context.Context, task *Task) error
	Delete(c context.Context, id string) error
}

type TaskUsecase interface {
	Create(c context.Context, task *Task, userID string) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
	Update(c context.Context, task *Task) error
	Delete(c context.Context, id string) error
}
