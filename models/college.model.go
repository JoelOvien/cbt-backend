package models

import (
	"time"
)

// College table
type College struct {
	CollegeID        string    `gorm:"size:6;primaryKey" json:"CollegeID"`
	CollegeLongName  string    `gorm:"size:50" json:"CollegeLongName"`
	CollegeShortName string    `gorm:"size:50" json:"CollegeShortName"`
	CurrentDean      string    `gorm:"size:15" json:"CurrentDean"`
	CollegeStatus    int8      `gorm:"type:tinyint" json:"CollegeStatus"`
	DateCreated      time.Time `gorm:"type:date" json:"DateCreated"`
	DateUpdated      time.Time `gorm:"type:date" json:"DateUpdated"`
}
