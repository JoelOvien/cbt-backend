package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Users struct for all users
type Users struct {
	UserID         string    `gorm:"column:UserID;not null" json:"UserID" validate:"required"`
	Surname        string    `gorm:"column:Surname;not null" json:"Surname" validate:"required"`
	FirstName      string    `gorm:"column:FirstName;not null" json:"FirstName" validate:"required"`
	EmailAddress   string    `gorm:"column:EmailAddress;not null" json:"EmailAddress" validate:"email,required"`
	Password       string    `gorm:"column:Password;not null" json:"Password" validate:"required,min=8"`
	UserStatus     string    `gorm:"column:UserStatus;not null" json:"UserStatus" validate:"required"`
	UserType       string    `gorm:"column:UserType;not null" json:"UserType" validate:"required"`
	DateCreated    time.Time `gorm:"column:DateCreated;not null" json:"DateCreated"`
	DateUpdated    time.Time `gorm:"column:DateUpdated;not null" json:"DateUpdated"`
	LastAccessDate time.Time `gorm:"column:LastAccessDate;not null" json:"LastAccessDate"`
	DepartmentID   string    `gorm:"column:DepartmentID;not null" json:"DepartmentID" validate:"required"`
	Role           string    `gorm:"column:Role;not null" json:"Role" validate:"required"`
}

var validate = validator.New()

// ErrorResponse struct for error response
type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// ValidateStruct validates the struct passed to it
func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// UserResponse struct for user response
type UserResponse struct {
	UserID         string    `gorm:"not null" json:"UserID"`
	Surname        string    `gorm:"not null" json:"Surname"`
	FirstName      string    `gorm:"not null" json:"FirstName"`
	EmailAddress   string    `gorm:"not null" json:"EmailAddress" validate:"email,required"`
	UserStatus     string    `gorm:"not null" json:"UserStatus"`
	UserType       string    `gorm:"not null" json:"UserType"`
	DateCreated    time.Time `gorm:"not null" json:"DateCreated"`
	DateUpdated    time.Time `gorm:"not null" json:"DateUpdated"`
	LastAccessDate time.Time `gorm:"not null" json:"LastAccessDate"`
	DepartmentID   string    `gorm:"not null" json:"DepartmentID"`
	Role           string    `gorm:"not null" json:"Role"`
}
