package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

type CreateUserResp struct {
	Id string `json:"id"`
}

type Error struct {
	Message string `json:"message"`
}
