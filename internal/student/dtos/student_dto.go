package dtos

import (
	"time"
)

type StudentDTO struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"student_id"`
	FullName  string    `gorm:"type:varchar(255);not null" json:"full_name"`
	BirthDate string    `gorm:"type:date;not null" json:"birth_date"`
	Gender    string    `gorm:"not null" json:"gender"`
	Faculty   int       `gorm:"not null" json:"faculty_id"`
	Course    int       `gorm:"not null" json:"course_id"`
	Program   int       `gorm:"not null" json:"program_id"`
	Address   string    `gorm:"type:text" json:"address,omitempty"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`
	Status    int       `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
