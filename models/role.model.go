package models

import (
	"github.com/google/uuid"
	"time"
)

// Role table struct
type Role struct {
	RoleID      uuid.UUID `gorm:"column:RoleID;type:uuid;default:uuid_generate_v4();primary_key" json:"RoleID"`
	RoleName    string    `gorm:"column:RoleName;size:10" json:"RoleName" validate:"required"`
	UserID      string    `gorm:"column:UserID;size:15" json:"UserID" validate:"required"`
	DateCreated time.Time `gorm:"column:DateCreated;type:date" json:"DateCreated"`
	DateUpdated time.Time `gorm:"column:DateUpdated;type:date" json:"DateUpdated"`
}
