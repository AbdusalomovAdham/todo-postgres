package models

import "time"

type Task struct {
	Uid        string    `gorm:"primaryKey;not null" json:"uid"`
	Done       bool      `gorm:"not null" json:"done"`
	Task       string    `gorm:"not null" json:"task"`
	Created_by string    `gorm:"not null" json:"created_by"`
	Created_at time.Time `gorm:"autoCreateTime" json:"created_at"`
	User       User      `gorm:"foreignKey:Created_by;references:Uid" json:"user"`
}
