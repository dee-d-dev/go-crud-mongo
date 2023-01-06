package models

import "go.mongodb.org/mongo-driver/mongo/bson/primitive"

type Todo struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`,
	Title string `json:"title"`
	Completed bool `json:"completed"`
}