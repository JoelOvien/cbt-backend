package models

import (
	"time"
)

// Result table
type Result struct {
	Session     string    `gorm:"size:9;primaryKey" json:"Session" validate:"required"`
	Semester    string    `gorm:"size:1;primaryKey" json:"Semester" validate:"required"`
	UserID      string    `gorm:"size:15;primaryKey" json:"UserID" validate:"required"`
	CourseCode  string    `gorm:"size:6" json:"CourseCode" validate:"required"`
	FinalScore  int       `gorm:"type:int" json:"FinalScore" validate:"required"`
	DateCreated time.Time `gorm:"type:date" json:"DateCreated" validate:"required"`
	DateUpdated time.Time `gorm:"type:date" json:"DateUpdated" validate:"required"`
}
