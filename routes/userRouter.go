package routes

import (
	"context"
	"log"
	"myproject/controllers"
	"myproject/repositories"
	"myproject/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserRouter(r *gin.Engine) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("todo-app")
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// create user
	r.POST("/", userController.CreateUser)

	// get user
	r.POST("/user", userController.GetUser)
}
