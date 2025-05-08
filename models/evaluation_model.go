package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EvaluationModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Value   string             `bson:"value" json:"value"`
	Date    string             `bson:"date" json:"date"`
	Student primitive.ObjectID `bson:"student" json:"student"`
	Object  primitive.ObjectID `bson:"object" json:"object"`
}

type EvaluationModelWithStudent struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Value   string             `bson:"value" json:"value"`
	Date    string             `bson:"date" json:"date"`
	Student string             `bson:"student" json:"student"`
	Object  string             `bson:"object" json:"object"`
}
