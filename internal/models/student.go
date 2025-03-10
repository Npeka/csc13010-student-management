package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Student struct with data validation
type Student struct {
	ID             uint       `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty" csv:"id,omitempty"`
	UserID         *uuid.UUID `gorm:"uniqueIndex" json:"user_id,omitempty"`
	User           User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	StudentID      string     `gorm:"type:varchar(12);uniqueIndex;not null" json:"student_id,omitempty" csv:"student_id,omitempty"`
	FullName       string     `gorm:"type:varchar(100);not null" json:"full_name,omitempty" csv:"full_name,omitempty"`
	BirthDate      string     `gorm:"type:date;not null" json:"birth_date,omitempty" csv:"birth_date,omitempty"`
	GenderID       uint       `gorm:"not null" json:"gender_id,omitempty" csv:"gender_id,omitempty"`
	Gender         Gender     `gorm:"foreignKey:GenderID" json:"gender,omitempty" csv:"gender,omitempty"`
	FacultyID      uint       `gorm:"not null" json:"faculty_id,omitempty" csv:"faculty_id,omitempty"`
	Faculty        Faculty    `gorm:"foreignKey:FacultyID" json:"faculty,omitempty" csv:"faculty,omitempty"`
	CourseID       uint       `gorm:"not null" json:"course_id,omitempty" csv:"course_id,omitempty"`
	Course         Course     `gorm:"foreignKey:CourseID" json:"course,omitempty" csv:"course,omitempty"`
	ProgramID      uint       `gorm:"not null" json:"program_id,omitempty" csv:"program_id,omitempty"`
	Program        Program    `gorm:"foreignKey:ProgramID" json:"program,omitempty" csv:"program,omitempty"`
	Address        string     `gorm:"type:text" json:"address,omitempty" csv:"address,omitempty"`
	Email          string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email,omitempty" csv:"email,omitempty"`
	Phone          string     `gorm:"type:varchar(20);not null" json:"phone,omitempty" csv:"phone,omitempty"`
	StatusID       uint       `gorm:"not null" json:"status_id,omitempty" csv:"status_id,omitempty"`
	Status         Status     `gorm:"foreignKey:StatusID" json:"status,omitempty" csv:"status,omitempty"`
	IsNotifyStatus bool       `gorm:"not null;default:false" json:"is_notify_status,omitempty" csv:"is_notify_status,omitempty"`
	CreatedAt      time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at,omitempty" csv:"created_at,omitempty"`
	UpdatedAt      time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at,omitempty" csv:"updated_at,omitempty"`
}

func (Student) TableName() string {
	return "students"
}

func (s *Student) BeforeDelete(tx *gorm.DB) error {
	var config Config
	err := tx.Where("id = 1").First(&config).Error
	if err != nil {
		return err
	}

	if config.DeleteLimit {
		return errors.New("deleting student is not allowed")
	}

	return nil
}

func (s *Student) BeforeSave(tx *gorm.DB) error {
	var config Config
	err := tx.Where("id = 1").First(&config).Error
	if err != nil {
		return err
	}

	// Validate StudentID (only alphanumeric, length 6-12 characters)
	if !isValidStudentID(s.StudentID) {
		return errors.New("invalid student_id, must be 6-12 alphanumeric characters")
	}

	// Validate FullName (only letters and spaces, max length 100 characters)
	if !isValidFullName(s.FullName) {
		return errors.New("invalid full_name, must contain only letters and spaces, max length 100")
	}

	// Validate BirthDate (must be before the current date)
	// string to time.Time
	_, err = time.Parse("2006-01-02", s.BirthDate)
	if err != nil {
		return errors.New("invalid birth_date, must be in the format YYYY-MM-DD")
	}

	// Validate Email
	if config.EmailDomain && !isValidEmail(s.Email) {
		return errors.New("invalid email, must be in the format @student.university.edu.vn")
	}

	// Validate Phone
	if config.ValidatePhone && !isValidPhone(s.Phone) {
		return errors.New("invalid phone number")
	}

	if config.StatusRules && s.ID > 0 {
		if oldStatusID, ok := tx.Statement.Context.Value("oldStatusID").(uint); ok {
			valid, err := IsValidStatusTransition(tx, oldStatusID, s.StatusID)
			if err != nil {
				return err
			}
			if !valid {
				return errors.New("invalid student status transition for your role")
			}
		}
	}

	return nil
}

func isValidStudentID(studentID string) bool {
	regex := `^[a-zA-Z0-9]{6,12}$`
	match, _ := regexp.MatchString(regex, studentID)
	return match
}

func isValidFullName(fullName string) bool {
	regex := `^[\p{L} ]{1,100}$`
	match, _ := regexp.MatchString(regex, fullName)
	return match
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@student\.university\.edu\.vn$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func isValidPhone(phone string) bool {
	regex := `^(\+84|0[3|5|7|8|9])\d{8}$`
	match, _ := regexp.MatchString(regex, phone)
	return match
}

func IsValidStatusTransition(tx *gorm.DB, currentStatusID, newStatusID uint) (bool, error) {
	if currentStatusID == newStatusID {
		return true, nil
	}

	var transition StatusTransition
	err := tx.Where("current_status_id = ? AND new_status_id = ?",
		currentStatusID, newStatusID).First(&transition).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
