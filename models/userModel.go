package models

type User struct {
	Uid      string `gorm:"primaryKey;not null"`
	Username string `gorm:"size:20;not null" json:"username" validate:"required,min=3,max=20"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
}
