package routes

import (
	"myproject/config"
	"myproject/controllers"
	"myproject/middleware"
	"myproject/repositories"
	"myproject/services"

	"github.com/gin-gonic/gin"
)

func TaskRouter(r *gin.Engine) {

	db := config.ConnectDatabase()
	
	repo := repositories.NewTaskRepository(db)
	service := services.NewTaskService(repo)
	controllers := controllers.NewTaskController(service)

	// get tasks
	r.GET("/tasks", middleware.AuthMiddleware(), controllers.GetTasks)

	// create task
	r.POST("/tasks", middleware.AuthMiddleware(), controllers.CreateTask)

	// update task
	r.PATCH("/:uid", middleware.AuthMiddleware(), controllers.UpdateTask)

	// delete task
	r.DELETE("/:uid", middleware.AuthMiddleware(), controllers.DeleteTask)
}
