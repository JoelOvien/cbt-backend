package models

import (
	"time"
)

// RegisteredCourses table struct
type RegisteredCourses struct {
	Session     string    `gorm:"column:Session;size:9;primaryKey" json:"Session"`
	Semester    string    `gorm:"column:Semester;size:1;primaryKey" json:"Semester"`
	UserID      string    `gorm:"column:UserID;not null" json:"UserID" validate:"required"`
	CourseCode  string    `gorm:"column:CourseCode;not null" json:"CourseCode" validate:"required"`
	DateCreated time.Time `gorm:"column:DateCreated;type:date" json:"DateCreated"`
	DateUpdated time.Time `gorm:"column:DateUpdated;type:date" json:"DateUpdated"`
}
