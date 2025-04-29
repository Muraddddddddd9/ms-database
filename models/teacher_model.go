package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TeachersModel represents a teachers document in MongoDB
// Uses primitive.ObjectID for status reference (likely a separate collection)
type TeachersModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // MongoDB ID
	Name       string             `bson:"name" json:"name"`                   // Teacher's first name
	Surname    string             `bson:"surname" json:"surname"`             // Teacher's last name
	Patronymic string             `bson:"patronymic" json:"patronymic"`       // Middle name (optional)
	Email      string             `bson:"email" json:"email"`                 // Contact email
	Password   string             `bson:"password" json:"password"`           // Auth password (should be hashed)
	Telegram   string             `bson:"telegram" json:"telegram"`           // Telegram username
	IPs        []string           `bson:"ips" json:"ips"`                     // Allowed IP addresses
	Status     primitive.ObjectID `bson:"status" json:"status"`               // Reference to status document
}

// TeachersWithStatusModel is similar to TeachersModel but with embedded status (string instead of reference)
type TeachersWithStatusModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Surname    string             `bson:"surname" json:"surname"`
	Patronymic string             `bson:"patronymic" json:"patronymic"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password"`
	Telegram   string             `bson:"telegram" json:"telegram"`
	IPs        []string           `bson:"ips" json:"ips"`
	Status     string             `bson:"status" json:"status"` // Direct status string instead of reference
}
