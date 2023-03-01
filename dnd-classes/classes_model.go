package dndclasses

import "go.mongodb.org/mongo-driver/bson/primitive"

type Classes struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Key  string             `json:"key,omitempty" validate:"required"`
	Name string             `json:"name,omitempty" validate:"required"`
	URL  string             `json:"url,omitempty" validate:"required"`
}
