package models

import (
	"time"

	"github.com/google/uuid"
)

// ExamTimetable table
type ExamTimetable struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"ID"`
	Session       string    `gorm:"column:Session;size:9" json:"Session" validate:"required"`
	Semester      string    `gorm:"column:Semester;size:1" json:"Semester" validate:"required"`
	CourseCode    string    `gorm:"column:CourseCode;size:6" json:"CourseCode" validate:"required"`
	ExamDate      time.Time `gorm:"column:ExamDate;type:date" json:"ExamDate" validate:"required"`
	ExamStartTime time.Time `gorm:"column:ExamStartTime;type:date" json:"ExamStartTime" validate:"required"`
	ExamEndTime   time.Time `gorm:"column:ExamEndTime;type:date" json:"ExamEndTime" validate:"required"`
	DisallowTime  int       `gorm:"column:DisallowTime;type:int" json:"DisallowTime" validate:"required"`
	ChiefExaminer string    `gorm:"column:ChiefExaminer;size:15" json:"ChiefExaminer" validate:"required"`
	Examiner1     string    `gorm:"column:Examiner1;size:15" json:"Examiner1" validate:"required"`
	Examiner2     string    `gorm:"column:Examiner2;size:15" json:"Examiner2"`
	Examiner3     string    `gorm:"column:Examiner3;size:15" json:"Examiner3"`
	Examiner4     string    `gorm:"column:Examiner4;size:15" json:"Examiner4"`
	Examiner5     string    `gorm:"column:Examiner5;size:15" json:"Examiner5"`
	Examiner6     string    `gorm:"column:Examiner6;size:15" json:"Examiner6"`
	Examiner7     string    `gorm:"column:Examiner7;size:15" json:"Examiner7"`
	Examiner8     string    `gorm:"column:Examiner8;size:15" json:"Examiner8"`
	Examiner9     string    `gorm:"column:Examiner9;size:15" json:"Examiner9"`
	Examiner10    string    `gorm:"column:Examiner10;size:15" json:"Examiner10"`
	DateCreated   time.Time `gorm:"column:DateCreated;type:date"`
	DateUpdated   time.Time `gorm:"column:DateUpdated;type:date"`
}
