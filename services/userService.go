package services

import (
	"encoding/json"
	"errors"

	"myproject/jwt"
	"myproject/models"
	"myproject/repositories"
	"myproject/utils"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

var (
	ErrHashFailed      = errors.New("could not hash password")
	ErrTokenGeneration = errors.New("error while generation token")
)

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{repo: userRepo}
}

// hash password func
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (us *UserService) CreateUser(user models.User) (string, error) {
	// hash password
	hashedPwd, err := HashPassword(user.Password)
	if err != nil {
		return "", ErrHashFailed
	}

	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Uid:      uuid.New().String(),
		Username: user.Username,
		Password: hashedPwd,
		Email:    user.Email,
	}

	// send repository
	if err := us.repo.Create(newUser); err != nil {
		return "", err
	}

	// sign token
	token, err := jwt.SignToken(newUser.Uid, user.Username)
	if err != nil {
		return "", ErrTokenGeneration
	}

	// marshal user for redis
	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		return "", err
	}

	// set redis
	err = utils.Set("user:"+user.Uid, jsonUser, 5*time.Minute)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *UserService) GetUser(user models.User) (models.User, string, error) {
	getUser, err := us.repo.GetUser(user.Username, user.Password)
	if err != nil {
		return models.User{}, "", err
	}

	getUser.Password = ""
	jsonUser, err := json.Marshal(getUser)
	if err != nil {
		return models.User{}, "", err
	}

	// create redis
	err = utils.Set("user:"+getUser.Uid, jsonUser, 10*time.Minute)
	if err != nil {
		return models.User{}, "", err
	}

	// sign token
	token, err := jwt.SignToken(getUser.Uid, getUser.Username)
	if err != nil {
		return models.User{}, "", err
	}

	return getUser, token, nil
}
