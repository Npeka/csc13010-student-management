package models

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

// Student struct with data validation
type Student struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	StudentID string    `gorm:"type:varchar(12);uniqueIndex;not null" json:"student_id"`
	FullName  string    `gorm:"type:varchar(100);not null" json:"full_name"`
	BirthDate time.Time `gorm:"type:date;not null" json:"birth_date"`
	GenderID  int       `gorm:"not null" json:"gender_id"`
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

// TableName - table name in DB
func (Student) TableName() string {
	return "students"
}

// Validate data before saving
func (s *Student) BeforeSave(tx *gorm.DB) (err error) {
	// Validate StudentID (only alphanumeric, length 6-12 characters)
	if !isValidStudentID(s.StudentID) {
		return errors.New("invalid student_id, must be 6-12 alphanumeric characters")
	}

	// Validate FullName (only letters and spaces, max length 100 characters)
	if !isValidFullName(s.FullName) {
		return errors.New("invalid full_name, must contain only letters and spaces, max length 100")
	}

	// Validate BirthDate (must be before the current date)
	if s.BirthDate.After(time.Now()) {
		return errors.New("invalid birth_date, cannot be in the future")
	}

	// Validate Email
	if !isValidEmail(s.Email) {
		return errors.New("invalid email, must be in the format @student.university.edu.vn")
	}

	// Validate Phone
	if !isValidPhone(s.Phone) {
		return errors.New("invalid phone number")
	}

	// Validate valid status transition
	if !isValidStatusTransition(tx, s.ID, s.StudentID, s.StatusID) {
		return errors.New("invalid student status transition")
	}

	return nil
}

// Function to validate StudentID
func isValidStudentID(studentID string) bool {
	regex := `^[a-zA-Z0-9]{6,12}$`
	match, _ := regexp.MatchString(regex, studentID)
	return match
}

// Function to validate FullName
func isValidFullName(fullName string) bool {
	regex := `^[\p{L} ]{1,100}$`
	match, _ := regexp.MatchString(regex, fullName)
	return match
}

// Function to validate Email
func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@student\.university\.edu\.vn$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

// Function to validate Phone (valid Vietnamese phone number)
func isValidPhone(phone string) bool {
	regex := `^(\+84|0[3|5|7|8|9])\d{8}$`
	match, _ := regexp.MatchString(regex, phone)
	return match
}

// Function to check valid status transition
func isValidStatusTransition(tx *gorm.DB, id int, studentID string, newStatusID int) bool {
	if id == 0 {
		// If new student, no need to check
		return true
	}

	var currentStudent Student
	if err := tx.Where("student_id = ?", studentID).First(&currentStudent).Error; err != nil {
		return false
	}

	// Valid status transition rules
	validTransitions := map[int][]int{
		1: {1, 2, 3, 4}, // "Đang học" -> "Tốt nghiệp", "Bỏ học", "Đình chỉ"
		2: {2},          // "Đã tốt nghiệp" -> KHÔNG thể thay đổi
		3: {3},          // "Đã bỏ học" -> KHÔNG thể thay đổi
		4: {4},          // "Bị đình chỉ" -> KHÔNG thể thay đổi
	}

	// Check if new status is valid
	for _, valid := range validTransitions[currentStudent.StatusID] {
		if valid == newStatusID {
			return true
		}
	}
	return false
}

// Log changes after creating student
func (s *Student) AfterCreate(tx *gorm.DB) (err error) {
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

// Log changes after updating student
func (s *Student) AfterUpdate(tx *gorm.DB) (err error) {
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
