package services

import (
	"encoding/json"
	"myproject/jwt"
	"myproject/models"
	"myproject/repositories"
	"myproject/utils"
	"time"

	"github.com/google/uuid"
)

type TaskService struct {
	repo *repositories.TaskRepostiry
}

func NewTaskService(repo *repositories.TaskRepostiry) *TaskService {
	return &TaskService{repo: repo}
}

func (ts *TaskService) GetTasks(authorization string) ([]models.Task, error) {
	// parse token
	parseToken, err := jwt.ParseToken(authorization)
	if err != nil {
		return nil, err
	}

	cacheKay := "tasks:" + parseToken.Uid

	cached, err := utils.Get(cacheKay)

	if err == nil && cached != "" {
		var tasks []models.Task
		if err := json.Unmarshal([]byte(cached), &tasks); err == nil {
			return tasks, nil
		}
	}

	// get tasks
	tasks, err := ts.repo.Get(parseToken.Uid)
	if err != nil {
		return nil, err
	}

	// write redis
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}
	err = utils.Set(cacheKay, jsonTasks, 10*time.Minute)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskService) CreateTask(taskText, authorization string) error {
	// parse token
	parseToken, err := jwt.ParseToken(authorization)
	if err != nil {
		return err
	}

	// create a new task
	task := models.Task{
		Uid:        uuid.New().String(),
		Created_by: parseToken.Uid,
		Done:       false,
		Task:       taskText,
		Created_at: time.Now(),
	}

	// delete tasks from redis
	_ = utils.Delete("tasks:" + parseToken.Uid)

	if err = ts.repo.Create(task); err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) UpdateTask(task, authorization, uid string) error {
	// parse token
	parseToken, err := jwt.ParseToken(authorization)
	if err != nil {
		return err
	}
	// update task
	if err := ts.repo.Update(task, uid); err != nil {
		return err
	}

	// delete tasks from redis
	_ = utils.Delete("tasks:" + parseToken.Uid)
	return nil
}

func (ts *TaskService) DeleteTask(uid, authorization string) error {
	// parse token
	parseToken, err := jwt.ParseToken(authorization)
	if err != nil {
		return err
	}

	// delete task
	if err := ts.repo.Delete(uid); err != nil {
		return err
	}

	//delete tasks from redis
	_ = utils.Delete("tasks:" + parseToken.Uid)
	return nil
}
