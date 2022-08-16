package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"Name,omitempty"`
	Phone string             `bson:"Phone"`
}
