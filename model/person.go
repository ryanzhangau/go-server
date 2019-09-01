package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Person - person struct
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// Persons - array of persons
type Persons []Person
