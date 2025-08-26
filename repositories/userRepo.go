package repositories

import (
	"context"
	"errors"
	"myproject/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CheckHashPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

type UserRepository struct {
	collections *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collections: db.Collection("users"),
	}
}

var (
	ErrUsernameTaken = errors.New("username already taken")
)

func (r *UserRepository) Create(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// chack username
	var existing models.User
	err := r.collections.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existing)

	if err == nil {
		return ErrUsernameTaken
	}

	if err != mongo.ErrNoDocuments {
		return err
	}

	// create uid
	user.Uid = uuid.New().String()

	_, err = r.collections.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUser(username, password string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	// get user
	err := r.collections.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	if !CheckHashPassword(password, user.Password) {
		return models.User{}, errors.New("incorrect password")
	}

	return user, nil
}
