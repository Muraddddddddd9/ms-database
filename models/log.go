package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	API    string             `bson:"api" json:"api"`
	Method string             `bson:"method" json:"method"`
	Status string             `bson:"status" json:"status"`
	Data   any                `bson:"data" json:"data"`
	Date   string             `bson:"date" json:"date"`
	Error  any                `bson:"error" json:"error"`
}
