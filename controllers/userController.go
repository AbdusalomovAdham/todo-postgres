package controllers

import (
	"log"
	"myproject/models"
	"myproject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	// get body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println("user asfdasdfasf", user)

	// validate user
	if errs := ValidateStruct(user); errs != nil {
		c.JSON(400, gin.H{"error": errs})
		return
	}

	// create user to database
	result, err := uc.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": result,
	})
}

func (uc *UserController) GetUser(c *gin.Context) {
	var user models.User

	// get body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get user and token
	user, token, err := uc.service.GetUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})

}
