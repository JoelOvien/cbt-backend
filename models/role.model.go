package models

import (
	"time"
)

// Role table struct
type Role struct {
	RoleID      int       `gorm:"primaryKey" json:"RoleID"`
	RoleName    string    `gorm:"size:10" json:"RoleName"`
	UserID      string    `gorm:"size:15" json:"UserID"`
	DateCreated time.Time `gorm:"type:date" json:"DateCreated"`
	DateUpdated time.Time `gorm:"type:date" json:"DateUpdated"`
}
