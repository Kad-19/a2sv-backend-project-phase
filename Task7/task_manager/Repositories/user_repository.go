package repository

import (
	"context"
	"task_manager/Domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) Domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *UserRepository) Create(c context.Context, user *Domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)
	return err
}

func (ur *UserRepository) FetchByEmail(c context.Context, email string) (*Domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	filter := bson.D{{Key: "email", Value: email}}

	var user Domain.User
	err := collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
