package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ObjectsGroupsModel represents a many-to-many relationship between:
// - Subjects (Objects)
// - Student Groups
// - Teachers
// This is the normalized version with ObjectID references
type ObjectsGroupsModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // Junction table ID
	Object  primitive.ObjectID `bson:"object" json:"object"`               // Reference to Subjects (ObjectsModel)
	Group   primitive.ObjectID `bson:"group" json:"group"`                 // Reference to Student Group
	Teacher primitive.ObjectID `bson:"teacher" json:"teacher"`             // Reference to Teacher
}

// ObjectsGroupsWithGroupAndTeacherModel is a denormalized version that includes:
// - Subject name (instead of reference)
// - Group name (instead of reference) 
// - Teacher name (instead of reference)
// Used for read operations where joined data is needed
type ObjectsGroupsWithGroupAndTeacherModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Object  string             `bson:"object" json:"object"`   // Subject name
	Group   string             `bson:"group" json:"group"`     // Group name
	Teacher string             `bson:"teacher" json:"teacher"` // Teacher's full name
}