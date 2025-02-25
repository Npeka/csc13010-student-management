package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Gender int

const (
	GenderMale Gender = iota
	GenderFemale
)

const (
	GenderMaleString   string = "male"
	GenderFemaleString string = "female"
)

var GenderMap = map[string]Gender{
	GenderMaleString:   GenderMale,
	GenderFemaleString: GenderFemale,
}

type Student struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	StudentID string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"student_id"`
	FullName  string    `gorm:"type:varchar(255);not null" json:"full_name"`
	BirthDate time.Time `gorm:"type:date;not null" json:"birth_date"`
	Gender    Gender    `gorm:"not null" json:"gender"`
	FacultyID int       `gorm:"not null" json:"faculty_id"`
	CourseID  int       `gorm:"not null" json:"course_id"`
	ProgramID int       `gorm:"not null" json:"program_id"`
	Address   string    `gorm:"type:text" json:"address,omitempty"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`
	StatusID  int       `gorm:"not null" json:"status_id"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Student) TableName() string {
	return "students"
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	newJSON, err := json.Marshal(s)
	if err != nil {
		return err
	}

	auditLog := AuditLog{
		TableName:     s.TableName(),
		RecordID:      s.ID,
		Action:        ActionCreate,
		ChangedFields: string(newJSON),
		ChangedBy:     RoleAdmin,
		CreatedAt:     time.Now(),
	}

	if err := tx.Model(&AuditLog{}).Create(&auditLog).Error; err != nil {
		return err
	}

	return nil
}

func (s *Student) BeforeUpdate(tx *gorm.DB) (err error) {
	changedJSON, err := json.Marshal(s)
	if err != nil {
		return err
	}

	auditLog := AuditLog{
		TableName:     s.TableName(),
		RecordID:      s.ID,
		Action:        ActionUpdate,
		ChangedFields: string(changedJSON),
		ChangedBy:     RoleAdmin,
		CreatedAt:     time.Now(),
	}

	if err := tx.Model(&AuditLog{}).Create(&auditLog).Error; err != nil {
		return err
	}

	return nil
}
