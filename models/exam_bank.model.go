package models

import (
	"github.com/google/uuid"
	"time"
)

// ExamBank table struc
type ExamBank struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"ID"`
	UserID         string    `gorm:"size:15" json:"UserID"`
	QuestionID     int16     `gorm:"type:smallint" json:"QuestionID"`
	CorrectAnswer  string    `gorm:"type:text" json:"CorrectAnswer"`
	AnswerProvided string    `gorm:"type:text" json:"AnswerProvided" validate:"required"`
	// AnswerMark could be correct or wrong (0/1)
	AnswerMark int8   `gorm:"type:tinyint" json:"AnswerMark" validate:"required"`
	Session    string `gorm:"size:9" json:"Session" validate:"required"`
	// Semester could be F or S for first and second
	Semester    string    `gorm:"size:1" json:"Semester" validate:"required"`
	DateCreated time.Time `gorm:"type:date" json:"DateCreated" validate:"required"`
	DateUpdated time.Time `gorm:"type:date" json:"DateUpdated" validate:"required"`
}
