package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetTasks returns the list of tasks
func GetTasks() ([]models.Task, error) {
	cur, err := tasks_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var tasks []models.Task
	for cur.Next(context.TODO()) {
		var task models.Task
		if err := cur.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// getTaskByID returns a task by its ID
func GetTaskByID(id string) *models.Task {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	var task models.Task
	err = tasks_collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil
	}
	return &task
}

// CreateTask adds a new task to the list
func CreateTask(task models.Task) (models.Task, error) {
	task.Status = models.Pending
	res, err := tasks_collection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.Task{}, err
	}

	// Set the ID field in the task to the inserted ObjectID's hex string
	if oid, ok := res.InsertedID.(bson.ObjectID); ok {
		task.ID = oid
	} else {
		return models.Task{}, errors.New("failed to retrieve inserted task")
	}

	return task, nil
}

// UpdateTask updates an existing task
func UpdateTask(id string, updatedTask models.Task) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	_, err = tasks_collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: updatedTask}})
	return err
}
	
// DeleteTask removes a task by its ID
func DeleteTask(id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	result, err := tasks_collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

// CreateUser adds a new user to the list
func CreateUser(user models.User) (models.User, error) {
	res, err := user_collection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, err
	}

	// Set the ID field in the user to the inserted ObjectID's hex string
	if oid, ok := res.InsertedID.(bson.ObjectID); ok {
		user.ID = oid
	} else {
		return models.User{}, errors.New("failed to retrieve inserted user")
	}

	return user, nil
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(email string) (*models.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	var user models.User
	err := user_collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}