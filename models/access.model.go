package models

import (
	"github.com/google/uuid"
	"time"
)

// Access table
type Access struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"ID"`
	UserID                 string    `gorm:"size:15" json:"UserID"`
	TransactionDescription string    `gorm:"type:text" json:"TransactionDescription"`
	TransactionDate        time.Time `gorm:"type:datetime" json:"TransactionDate"`
}
