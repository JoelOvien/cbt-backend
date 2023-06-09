package models

import (
	"time"
)

// Course table
type Course struct {
	CourseCode  string    `gorm:"column:CourseCode;not null" json:"CourseCode" validate:"required"`
	CourseTitle string    `gorm:"column:CourseTitle;not null" json:"CourseTitle" validate:"required"`
	CreditUnits int16     `gorm:"column:CreditUnits;not null" json:"CreditUnits" validate:"required"`
	DateCreated time.Time `gorm:"column:DateCreated;not null" json:"DateCreated"`
	DateUpdated time.Time `gorm:"column:DateUpdated;not null" json:"DateUpdated"`
}
