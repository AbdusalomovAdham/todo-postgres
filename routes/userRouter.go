package routes

import (

	"myproject/config"
	"myproject/controllers"
	"myproject/models"
	"myproject/repositories"
	"myproject/services"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	db := config.ConnectDatabase()

	db.AutoMigrate(&models.User{})
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// create user
	r.POST("/", userController.CreateUser)

	// get user
	r.POST("/user", userController.GetUser)
}
