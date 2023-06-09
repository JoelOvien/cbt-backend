package models

import (
	"time"
)

// College table
type College struct {
	CollegeID        string    `gorm:"column:CollegeID;size:6;primaryKey" json:"CollegeID"  validate:"required"`
	CollegeLongName  string    `gorm:"column:CollegeLongName;size:50" json:"CollegeLongName"  validate:"required" `
	CollegeShortName string    `gorm:"column:CollegeShortName;size:50" json:"CollegeShortName" validate:"required"`
	CurrentDean      string    `gorm:"column:CurrentDean;size:15" json:"CurrentDean" validate:"required" `
	CollegeStatus    int8      `gorm:"column:CollegeStatus;type:tinyint" json:"CollegeStatus" `
	DateCreated      time.Time `gorm:"column:DateCreated;type:date" json:"DateCreated"`
	DateUpdated      time.Time `gorm:"column:DateUpdated;type:date" json:"DateUpdated"`
}
