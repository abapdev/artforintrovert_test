package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"Name,omitempty"`
	Phone string             `bson:"Phone"`
}
type DataJSON struct {
	Name  string `json:"Name"`
	Phone string `json:"Phone"`
}
type DataState struct {
	Name  string `json:"name"`
	State string `json:"state"`
}
type CachedData map[string]string
