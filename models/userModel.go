package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Uid      string             `bson:"uid" json:"uid"`
	Username string             `bson:"username" json:"username" validate:"required,min=3,max=20"`
	Password string             `bson:"password" json:"password" validate:"required,min=6"`
	Email    string             `bson:"email" json:"email" validate:"required,email"`
}
