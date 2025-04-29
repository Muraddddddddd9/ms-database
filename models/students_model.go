package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// StudentsModel represents a students document in MongoDB
// Uses references (ObjectID) for group and status (normalized structure)
type StudentsModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // MongoDB ID
	Name       string             `bson:"name" json:"name"`                   // Student's first name
	Surname    string             `bson:"surname" json:"surname"`             // Student's last name
	Patronymic string             `bson:"patronymic" json:"patronymic"`       // Middle name (optional)
	Group      primitive.ObjectID `bson:"group" json:"group"`                 // Reference to group document
	Email      string             `bson:"email" json:"email"`                 // Contact email
	Password   string             `bson:"password" json:"password"`           // Hashed password (never plaintext)
	Telegram   string             `bson:"telegram" json:"telegram"`           // Telegram username
	Diplomas   []string           `bson:"diplomas" json:"diplomas"`           // List of diploma IDs/names
	IPs        []string           `bson:"ips" json:"ips"`                     // Allowed IP addresses
	Status     primitive.ObjectID `bson:"status" json:"status"`               // Reference to status document
}

// StudentsWithGroupAndStatusModel represents a student with embedded group and status data
// (Denormalized structure for easier queries/response handling)
type StudentsWithGroupAndStatusModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Surname    string             `bson:"surname" json:"surname"`
	Patronymic string             `bson:"patronymic" json:"patronymic"`
	Group      string             `bson:"group" json:"group"` // Group name (embedded)
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password"`
	Telegram   string             `bson:"telegram" json:"telegram"`
	Diplomas   []string           `bson:"diplomas" json:"diplomas"`
	IPs        []string           `bson:"ips" json:"ips"`
	Status     string             `bson:"status" json:"status"` // Status name (embedded)
}
