package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// User struct for all users
type User struct {
	UserID         string    `gorm:"not null" json:"UserID"`
	Surname        string    `gorm:"not null" json:"Surname"`
	FirstName      string    `gorm:"not null" json:"FirstName"`
	EmailAddress   string    `gorm:"not null" json:"EmailAddress" validate:"email,required"`
	Password       *string   `gorm:"not null" json:"Password" validate:"required,min=8"`
	UserStatus     string    `gorm:"not null" json:"UserStatus"`
	UserType       string    `gorm:"not null" json:"UserType"`
	DateCreated    time.Time `gorm:"not null" json:"DateCreated"`
	DateUpdated    time.Time `gorm:"not null" json:"DateUpdated"`
	LastAccessDate time.Time `gorm:"not null" json:"LastAccessDate"`
	DepartmentID   string    `gorm:"not null" json:"DepartmentID"`
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
}
