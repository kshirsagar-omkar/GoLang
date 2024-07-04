package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MoneyManagement struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tag		string			   `json:"tag,omitempty"`
	ForWhat string             `json:"forwhat,omitempty"`
	Money   string             `json:"money,omitempty"`
	Time 	string			   `json:"time,omitempty"`
}