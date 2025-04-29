package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// GroupsModel represents a student group in the system
// Contains the group name and reference to the assigned teacher
type GroupsModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // Unique group identifier
	Group   string             `bson:"group" json:"group"`                 // Group name
	Teacher primitive.ObjectID `bson:"teacher" json:"teacher"`             // Reference to the teacher's ID
}

// GroupsWithTeacherModel represents a group with embedded teacher name
// Used when you need to display group information including teacher name
type GroupsWithTeacherModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Group   string             `bson:"group" json:"group"`     // Group name
	Teacher string             `bson:"teacher" json:"teacher"` // Teacher's full name
}