package models

import (
	"time"
)

// Department table
type Department struct {
	DepartmentID        string    `gorm:"column:DepartmentID;size:15;primaryKey" json:"DepartmentID" validate:"required"`
	DepartmentLongName  string    `gorm:"column:DepartmentLongName;size:50" json:"DepartmentLongName" validate:"required"`
	DepartmentShortName string    `gorm:"column:DepartmentShortName;size:15" json:"DepartmentShortName" validate:"required"`
	DepartmentStatus    int8      `gorm:"column:DepartmentStatus;type:tinyint" json:"DepartmentStatus"`
	CurrentHodID        string    `gorm:"column:CurrentHodID;size:15" json:"CurrentHodID" validate:"required"`
	CollegeID           string    `gorm:"column:CollegeID;size:15" json:"CollegeID" validate:"required"`
	DateCreated         time.Time `gorm:"column:DateCreated;type:date" json:"DateCreated"`
	DateUpdated         time.Time `gorm:"column:DateUpdated;type:date" json:"DateUpdated"`
}
