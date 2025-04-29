package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SelectModels is a container for dropdown/select options
// Used to provide frontend with available choices for forms/filters
type SelectModels struct {
	Objects  []SelectObjectModels  `json:"select_object"`   // List of available subjects
	Groups   []SelectGroupModels   `json:"select_groups"`   // List of available groups
	Teachers []SelectTeacherModels `json:"select_teachers"` // List of available teachers
	Statuses []SelectStatusModels  `json:"select_statuses"` // List of available statuses
}

// SelectObjectModels represents a subject option for dropdowns
// Uses value/label pattern for frontend select components
type SelectObjectModels struct {
	ID     primitive.ObjectID `bson:"_id" json:"value"`    // Actual value (ObjectID)
	Object string             `bson:"object" json:"label"` // Display text
}

// SelectGroupModels represents a group option for dropdowns
type SelectGroupModels struct {
	ID    primitive.ObjectID `bson:"_id" json:"value"`   // Actual value (ObjectID)
	Group string             `bson:"group" json:"label"` // Display text
}

// SelectStatusModels represents a status option for dropdowns
type SelectStatusModels struct {
	ID     primitive.ObjectID `bson:"_id" json:"value"`    // Actual value (ObjectID)
	Status string             `bson:"status" json:"label"` // Display text
}

// SelectTeacherModels represents a teacher option for dropdowns
// Includes separate name fields for more flexible display options
type SelectTeacherModels struct {
	ID         primitive.ObjectID `bson:"_id" json:"value"`             // Actual value (ObjectID)
	Name       string             `bson:"name" json:"name"`             // First name
	Surname    string             `bson:"surname" json:"surname"`       // Last name
	Patronymic string             `bson:"patronymic" json:"patronymic"` // Middle name
}

// Log represents an API activity record
// Used for auditing and debugging purposes
type Log struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // Log entry ID
	API    string             `bson:"api" json:"api"`                     // API endpoint path
	Method string             `bson:"method" json:"method"`               // HTTP method
	Status string             `bson:"status" json:"status"`               // Result status
	Data   any                `bson:"data" json:"data"`                   // Request/response data
	Date   string             `bson:"date" json:"date"`                   // Timestamp of the event
	Error  any                `bson:"error" json:"error"`                 // Error details (if any)
}
