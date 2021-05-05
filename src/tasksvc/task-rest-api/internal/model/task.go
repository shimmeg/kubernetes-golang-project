package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" binding:"required" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
}
