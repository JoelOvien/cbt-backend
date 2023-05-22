package models

import (
	"github.com/google/uuid"
	"time"
)

// ExamTimetable table
type ExamTimetable struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"ID"`
	Session       string    `gorm:"size:9"`
	Semester      string    `gorm:"size:1"`
	CourseCode    string    `gorm:"size:6"`
	ExamDate      time.Time `gorm:"type:date"`
	ExamStartTime time.Time `gorm:"type:date"`
	ExamEndTime   time.Time `gorm:"type:date"`
	DisallowTime  int       `gorm:"type:int"`
	ChiefExaminer string    `gorm:"size:15"`
	Examiner1     string    `gorm:"size:15"`
	Examiner2     string    `gorm:"size:15"`
	Examiner3     string    `gorm:"size:15"`
	Examiner4     string    `gorm:"size:15"`
	Examiner5     string    `gorm:"size:15"`
	Examiner6     string    `gorm:"size:15"`
	Examiner7     string    `gorm:"size:15"`
	Examiner8     string    `gorm:"size:15"`
	Examiner9     string    `gorm:"size:15"`
	Examiner10    string    `gorm:"size:15"`
	DateCreated   time.Time `gorm:"type:date"`
	DateUpdated   time.Time `gorm:"type:date"`
}
