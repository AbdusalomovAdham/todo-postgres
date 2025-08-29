package repositories

import (
	"errors"

	"myproject/models"

	"gorm.io/gorm"
)

type TaskRepostiry struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepostiry {
	return &TaskRepostiry{db: db}
}

func (tr *TaskRepostiry) Get(uid string) ([]models.Task, error) {
	var tasks []models.Task

	// get tasks
	err := tr.db.
		Where("created_by = ?", uid).
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *TaskRepostiry) Create(task models.Task) error {

	// create task
	err := tr.db.Create(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepostiry) Update(newTask, task_uid string) error {
	// update task
	res := tr.db.Model(&models.Task{}).
		Where("uid = ?", task_uid).
		Updates(map[string]interface{}{
			"task": newTask,
		})

	if res.Error != nil {
		return errors.New("error updating task")
	}
	if res.RowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (tr *TaskRepostiry) Delete(uid string) error {
	
	// delete task
	res := tr.db.Where("uid = ?", uid).Delete(&models.Task{})
	if res.Error != nil {
		return errors.New("error deleting task")
	}
	if res.RowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil
}
