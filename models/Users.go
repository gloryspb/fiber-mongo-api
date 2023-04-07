package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Login    string             `json:"Login,omitempty" bson:"Login,omitempty"`
	Password string             `json:"Password,omitempty" bson:"Password,omitempty"`
}
