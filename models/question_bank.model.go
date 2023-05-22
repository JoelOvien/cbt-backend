package models

import (
	"time"
)

// QuestionsBank table struct
type QuestionsBank struct {
	QuestionID    int16     `gorm:"primaryKey" json:"QuestionID"`
	CourseCode    string    `gorm:"size:6" json:"CourseCode"`
	AnswerTypeID  string    `gorm:"size:2" json:"AnswerTypeID"`
	Question      string    `gorm:"type:text" json:"Question"`
	QuestionImage string    `gorm:"size:150" json:"QuestionImage"`
	Answer1       string    `gorm:"type:text" json:"Answer1"`
	Answer2       string    `gorm:"type:text" json:"Answer2"`
	Answer3       string    `gorm:"type:text" json:"Answer3"`
	Answer4       string    `gorm:"type:text" json:"Answer4"`
	Answer5       string    `gorm:"type:text" json:"Answer5"`
	CorrectAnswer string    `gorm:"type:text" json:"CorrectAnswer"`
	DateCreated   time.Time `gorm:"type:date" json:"DateCreated"`
	DateUpdated   time.Time `gorm:"type:date" json:"DateUpdated"`
}
