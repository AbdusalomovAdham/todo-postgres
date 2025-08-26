package repositories

import (
	"context"
	"myproject/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepostiry struct {
	collections *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) *TaskRepostiry {
	return &TaskRepostiry{collections: db.Collection("tasks")}
}

func (tr *TaskRepostiry) Get(uid string) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get tasks
	cursor, err := tr.collections.Find(ctx, bson.M{"created_by": uid})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// tasks get from cursor
	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *TaskRepostiry) Create(task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// insert task
	_, err := tr.collections.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepostiry) Update(newTask, task_uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// update task
	_, err := tr.collections.UpdateOne(ctx,
		bson.M{"uid": task_uid},
		bson.M{"$set": bson.M{"task": newTask}},
	)

	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepostiry) Delete(uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// delete task
	_, err := tr.collections.DeleteOne(ctx, bson.M{"uid": uid})
	if err != nil {
		return err
	}

	return nil
}
