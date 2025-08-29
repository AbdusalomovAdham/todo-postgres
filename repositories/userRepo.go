package repositories

import (
	"errors"
	"myproject/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckHashPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

var (
	ErrUsernameTaken = errors.New("username already taken")
)

func (r *UserRepository) Create(user models.User) error {
	var count int64

	// check user
	if err := r.db.Model(&models.User{}).
		Where("username = ?", user.Username).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return ErrUsernameTaken
	}

	// create user
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetUser(username, password string) (models.User, error) {

	var user models.User

	// get user
	err := r.db.
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return models.User{}, errors.New("password or username wrong")
	}

	// check password
	if !CheckHashPassword(password, user.Password) {
		return models.User{}, errors.New("password or username wrong")
	}

	return user, nil
}
