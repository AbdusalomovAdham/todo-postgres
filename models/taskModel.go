package models

import "time"

type Task struct {
	ID         string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Uid        string    `bson:"uid" json:"uid"`
	Created_by string    `bson:"created_by" json:"created_by"`
	Task       string    `bson:"task" json:"task"`
	Done       bool      `bson:"done" json:"done"`
	Created_at time.Time `bson:"created_at" json:"created_at"`
}
