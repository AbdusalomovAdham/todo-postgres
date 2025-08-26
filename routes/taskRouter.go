package routes

import (
	"context"
	"log"
	"myproject/controllers"
	"myproject/middleware"
	"myproject/repositories"
	"myproject/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TaskRouter(r *gin.Engine) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("todo-app")
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
