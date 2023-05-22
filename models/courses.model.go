package models

import (
	"time"
)

// Course table
type Course struct {
	CourseCode  string    `gorm:"size:6;primaryKey" json:"CourseCode"`
	CourseTitle string    `gorm:"size:20" json:"CourseTitle"`
	CreditUnits int16     `gorm:"type:smallint" json:"CreditUnits"`
	DateCreated time.Time `gorm:"type:date" json:"DateCreated"`
	DateUpdated time.Time `gorm:"type:date" json:"DateUpdated"`
}
