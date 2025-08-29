package controllers

import (
	"myproject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service *services.TaskService
}

type TaskInput struct {
	Task string `json:"task"`
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	// authorization get from header
	authorization := c.GetHeader("Authorization")

	// get task
	tasks, err := tc.service.GetTasks(authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	// get authorization
	authorization := c.GetHeader("Authorization")
	var input TaskInput

	// get body
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	// create task
	err := tc.service.CreateTask(input.Task, authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "creates task",
	})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	// get authorization
	authorization := c.GetHeader("Authorization")

	// get param
	uid := c.Param("uid")
	var input TaskInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	// update task
	if err := tc.service.UpdateTask(input.Task, authorization, uid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update task"})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	// get authorization
	authorization := c.GetHeader("Authorization")

	// get param uid
	uid := c.Param("uid")

	if err := tc.service.DeleteTask(uid, authorization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete task"})
}
