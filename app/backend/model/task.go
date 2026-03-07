package model

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Task struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Completed bool          `bson:"completed" json:"completed"`
}

type Request struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
