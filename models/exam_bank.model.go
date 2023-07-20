package models

import (
	"time"

	"github.com/google/uuid"
)

// ExamBank table struc
type ExamBank struct {
	ID             uuid.UUID `gorm:"column:ID;type:uuid;default:uuid_generate_v4();primary_key" json:"ID"`
	UserID         string    `gorm:"column:UserID;size:15" json:"UserID" validate:"required"`
	QuestionID     string    `gorm:"column:QuestionID;primaryKey;size:15" json:"QuestionID"`
	CourseCode     string    `gorm:"column:CourseCode;size:6" json:"CourseCode" validate:"required"`
	CorrectAnswer  string    `gorm:"column:CorrectAnswer;type:text" json:"CorrectAnswer" validate:"required"`
	AnswerProvided string    `gorm:"column:AnswerProvided;type:text" json:"AnswerProvided" validate:"required"`
	// AnswerMark could be correct or wrong (0/1)
	AnswerMark int8   `gorm:"column:AnswerMark;type:tinyint" json:"AnswerMark"`
	Session    string `gorm:"column:Session;size:9" json:"Session" validate:"required"`
	// Semester could be F or S for first and second
	Semester    string    `gorm:"column:Semester;size:1" json:"Semester" validate:"required"`
	DateCreated time.Time `gorm:"column:DateCreated;type:date"`
	DateUpdated time.Time `gorm:"column:DateUpdated;type:date"`
}
