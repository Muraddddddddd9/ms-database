package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContestsModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Teacher     primitive.ObjectID `bson:"teacher" json:"teacher"`
	Room        string             `bson:"room" json:"room"`
	Course      []string           `bson:"course" json:"course"`
	Start       string             `bson:"start" json:"start"`
	End         string             `bson:"end" json:"end"`
	Responsed   []ResponsedModel   `bson:"responsed" json:"responsed"`
}

type ResponsedModel struct {
	Students    primitive.ObjectID `bson:"student" json:"student"`
	Description string             `bson:"description" json:"description"`
}

type ContestsWithTeacherModel struct {
	ID          primitive.ObjectID          `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string                      `bson:"title" json:"title"`
	Description string                      `bson:"description" json:"description"`
	TeacherID   primitive.ObjectID          `bson:"teacher" json:"teacher"`
	TeacherName string                      `bson:"teacher_name" json:"teacher_name"`
	Room        string                      `bson:"room" json:"room"`
	Course      []string                    `bson:"course" json:"course"`
	Start       string                      `bson:"start" json:"start"`
	End         string                      `bson:"end" json:"end"`
	Responsed   []ResponsedWithStudentModel `bson:"responsed" json:"responsed"`
}

type ResponsedWithStudentModel struct {
	StudentName       string `bson:"name" json:"name"`
	StudentSurname    string `bson:"surname" json:"surname"`
	StudentPatronymic string `bson:"patronymic" json:"patronymic"`
	StudentGroup      string `bson:"group" json:"group"`
	StudentEmail      string `bson:"email" json:"email"`
	StudentTelegram   string `bson:"telegram" json:"telegram"`
	Description       string `bson:"description" json:"description"`
}
