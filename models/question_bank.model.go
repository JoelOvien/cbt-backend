package models

import (
	"time"
)

// QuestionsBank table struct
type QuestionsBank struct {
	QuestionNumber int       `gorm:"column:QuestionNumber;type:int" json:"QuestionNumber" validate:"required"`
	QuestionID     string    `gorm:"column:QuestionID;primaryKey;size:15" json:"QuestionID"`
	CourseCode     string    `gorm:"column:CourseCode;size:6" json:"CourseCode" validate:"required"`
	AnswerTypeID   string    `gorm:"column:AnswerTypeID;size:2" json:"AnswerTypeID" validate:"required"`
	Question       string    `gorm:"column:Question;type:text" json:"Question" validate:"required"`
	QuestionImage  string    `gorm:"column:QuestionImage;size:150" json:"QuestionImage" `
	Answer1        string    `gorm:"column:Answer1;type:text" json:"Answer1"`
	Answer2        string    `gorm:"column:Answer2;type:text" json:"Answer2"`
	Answer3        string    `gorm:"column:Answer3;type:text" json:"Answer3"`
	Answer4        string    `gorm:"column:Answer4;type:text" json:"Answer4"`
	Answer5        string    `gorm:"column:Answer5;type:text" json:"Answer5"`
	CorrectAnswer  string    `gorm:"column:CorrectAnswer;type:text" json:"CorrectAnswer" validate:"required"`
	DateCreated    time.Time `gorm:"column:DateCreated;type:date"`
	DateUpdated    time.Time `gorm:"column:DateUpdated;type:date"`
}
