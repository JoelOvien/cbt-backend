package models

import (
	"time"
)

// Department table
type Department struct {
	DepartmentID        string    `gorm:"size:15;primaryKey" json:"DepartmentID" validate:"required"`
	DepartmentLongName  string    `gorm:"size:50" json:"DepartmentLongName" validate:"required"`
	DepartmentShortName string    `gorm:"size:15" json:"DepartmentShortName" validate:"required"`
	DepartmentStatus    int8      `gorm:"type:tinyint" json:"DepartmentStatus" validate:"required"`
	CurrentHodID        string    `gorm:"size:15" json:"CurrentHodID" validate:"required"`
	CollegeID           string    `gorm:"size:15" json:"CollegeID" validate:"required"`
	DateCreated         time.Time `gorm:"type:date" json:"DateCreated" validate:"required"`
	DateUpdated         time.Time `gorm:"type:date" json:"DateUpdated" validate:"required"`
}
