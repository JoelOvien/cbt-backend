package models

import (
	"time"
)

// RegisteredCourses table struct
type RegisteredCourses struct {
	Session     string    `gorm:"size:9;primaryKey" json:"Session"`
	Semester    string    `gorm:"size:1;primaryKey" json:"Semester"`
	UserID      string    `gorm:"size:15" json:"UserID"`
	CourseCode  string    `gorm:"size:6" json:"CourseCode"`
	DateCreated time.Time `gorm:"type:date" json:"DateCreated"`
	DateUpdated time.Time `gorm:"type:date" json:"DateUpdated"`
}
