package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ObjectsModel represents academic subjects in the system
// Stores subject names like: "History", "Russian Language", "Mathematics", etc.
type ObjectsModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // Unique MongoDB identifier
	Object string             `bson:"object" json:"object"`               // Name of academic subject
}
