package models

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Student struct with data validation
type Student struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null" json:"id" csv:"id"`
	UserID    *uint     `gorm:"uniqueIndex" json:"user_id,omitempty"`
	User      *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	StudentID string    `gorm:"type:varchar(12);uniqueIndex;not null" json:"student_id" csv:"student_id"`
	FullName  string    `gorm:"type:varchar(100);not null" json:"full_name" csv:"full_name"`
	BirthDate string    `gorm:"type:date;not null" json:"birth_date" csv:"birth_date"`
	GenderID  uint      `gorm:"not null" json:"gender_id" csv:"gender_id"`
	Gender    Gender    `gorm:"foreignKey:GenderID" json:"gender,omitempty" csv:"gender"`
	FacultyID uint      `gorm:"not null" json:"faculty_id" csv:"faculty_id"`
	Faculty   Faculty   `gorm:"foreignKey:FacultyID" json:"faculty,omitempty" csv:"faculty"`
	CourseID  uint      `gorm:"not null" json:"course_id" csv:"course_id"`
	Course    Course    `gorm:"foreignKey:CourseID" json:"course,omitempty" csv:"course"`
	ProgramID uint      `gorm:"not null" json:"program_id" csv:"program_id"`
	Program   Program   `gorm:"foreignKey:ProgramID" json:"program,omitempty" csv:"program"`
	Address   string    `gorm:"type:text" json:"address,omitempty" csv:"address"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" csv:"email"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone" csv:"phone"`
	StatusID  uint      `gorm:"not null" json:"status_id" csv:"status_id"`
	Status    Status    `gorm:"foreignKey:StatusID" json:"status,omitempty" csv:"status"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at" csv:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at" csv:"updated_at"`
}

// TableName - table name in DB
func (Student) TableName() string {
	return "students"
}

// Validate data before saving
func (s *Student) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println(s)
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
	if !isValidEmail(s.Email) {
		return errors.New("invalid email, must be in the format @student.university.edu.vn")
	}

	// Validate Phone
	if !isValidPhone(s.Phone) {
		return errors.New("invalid phone number")
	}

	if s.ID > 0 {
		var currentStudent Student
		if err := tx.Select("status_id").Where("id = ?", s.ID).First(&currentStudent).Error; err != nil {
			return err
		}

		// Validate valid status transition
		userRole := tx.Statement.Context.Value("userRole").(string)

		// Check status transition
		valid, err := IsValidStatusTransition(tx, s.StatusID, s.StatusID, userRole)
		if err != nil {
			return err
		}
		if !valid {
			return errors.New("invalid student status transition for your role")
		}
	}

	// Set context for AfterCreate
	tx.Statement.Context = context.WithValue(tx.Statement.Context, "userData", struct {
		UserID uint
		Email  string
		Role   string
	}{
		// UserID: s.UserID,
		Email: s.Email,
		Role:  "student",
	})

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

func IsValidStatusTransition(tx *gorm.DB, currentStatusID, newStatusID uint, userRole string) (bool, error) {
	if currentStatusID == newStatusID {
		return true, nil
	}

	var transition StatusTransition
	err := tx.Where("current_status_id = ? AND new_status_id = ?",
		currentStatusID, newStatusID).First(&transition).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // Transition not found = not allowed
		}
		return false, err // Database error
	}

	// Check if user role is allowed for this transition
	allowedRoles := strings.Split(transition.AllowedRoles, ",")
	for _, role := range allowedRoles {
		if strings.TrimSpace(role) == userRole {
			return true, nil
		}
	}

	return false, nil // User role not allowed for this transition
}

// Log changes after creating student
func (s *Student) AfterCreate(tx *gorm.DB) (err error) {
	// Get user data from context
	userData := tx.Statement.Context.Value("userData").(struct {
		UserID uint
		Email  string
		Role   string
	})

	err = LogModelChanges(
		tx,
		s.TableName(),
		s.ID,
		ActionCreate,
		nil, // No old student for creation
		s,
		userData.UserID,
		userData.Email,
		"admin", // Hardcoded role for creation
	)

	return err
}

// Log changes after updating student
func (s *Student) AfterUpdate(tx *gorm.DB) (err error) {
	// Get the original student data before update
	var oldStudent Student
	if err := tx.Unscoped().Where("id = ?", s.ID).First(&oldStudent).Error; err != nil {
		return err
	}

	// Get user data from context
	userData := tx.Statement.Context.Value("userData").(struct {
		UserID uint
		Email  string
		Role   string
	})

	err = LogModelChanges(
		tx,
		s.TableName(),
		s.ID,
		ActionUpdate,
		oldStudent,
		s,
		userData.UserID,
		userData.Email,
		userData.Role,
	)

	return err
}

func (s *Student) AfterDelete(tx *gorm.DB) (err error) {
	// Get user data from context
	userData := tx.Statement.Context.Value("userData").(struct {
		UserID uint
		Email  string
		Role   string
	})

	err = LogModelChanges(
		tx,
		s.TableName(),
		s.ID,
		ActionDelete,
		s,   // The student data to be deleted
		nil, // No new student for deletion
		userData.UserID,
		userData.Email,
		userData.Role,
	)

	return err
}

// Function to check valid status transition
// func isValidStatusTransition(tx *gorm.DB, id int, studentID string, newStatusID int) bool {
// 	if id == 0 {
// 		// If new student, no need to check
// 		return true
// 	}

// 	var currentStudent Student
// 	if err := tx.Where("student_id = ?", studentID).First(&currentStudent).Error; err != nil {
// 		return false
// 	}

// 	// Valid status transition rules
// 	validTransitions := map[int][]int{
// 		1: {1, 2, 3, 4}, // "Đang học" -> "Tốt nghiệp", "Bỏ học", "Đình chỉ"
// 		2: {2},          // "Đã tốt nghiệp" -> KHÔNG thể thay đổi
// 		3: {3},          // "Đã bỏ học" -> KHÔNG thể thay đổi
// 		4: {4},          // "Bị đình chỉ" -> KHÔNG thể thay đổi
// 	}

// 	// Check if new status is valid
// 	for _, valid := range validTransitions[currentStudent.StatusID] {
// 		if valid == newStatusID {
// 			return true
// 		}
// 	}
// 	return false
// }
