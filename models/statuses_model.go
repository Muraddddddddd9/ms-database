package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// UserRoleModel defines user role types in the system
// Used to distinguish between: "student", "teacher", "admin"
type StatusesModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // Unique MongoDB identifier
	Status string             `bson:"status" json:"status"`               // Role type: "student"|"teacher"|"admin"
}
