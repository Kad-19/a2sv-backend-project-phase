package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string        `json:"email" binding:"required,email"`
	Password string        `json:"password" binding:"required,min=6"`
	Role     string        `json:"role" binding:"required,oneof=admin user"`
}