package repository

import (
	"context"
	"task_manager/Domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) Domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *Domain.Task, userId string) error {
	collection := tr.database.Collection(tr.collection)

	oid, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	task.UserID = oid

	_, err = collection.InsertOne(c, task)

	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]Domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	oid, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "user_id", Value: oid}}

	var tasks []Domain.Task
	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) Update(c context.Context, task *Domain.Task) error {
	collection := tr.database.Collection(tr.collection)

	filter := bson.D{{Key: "_id", Value: task.ID}}
	update := bson.D{{Key: "$set", Value: task}}

	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (tr *taskRepository) Delete(c context.Context, id string) error {
	collection := tr.database.Collection(tr.collection)

	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}

	_, err = collection.DeleteOne(c, filter)
	return err
}
