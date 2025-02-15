package tests

import (
	"bytes"
	"strings"
	"student-management/models"
	"student-management/services"
	"testing"
)

func TestIsStudentIDExists(t *testing.T) {
	services.Students = []models.Student{
		{ID: "123"},
		{ID: "456"},
	}

	tests := []struct {
		id       string
		expected bool
	}{
		{"789", false},
		{"123", true},
	}

	for _, test := range tests {
		result := services.IsStudentIDExists(test.id)
		if result != test.expected {
			t.Errorf("IsStudentIDExists(%s) = %v; expected %v", test.id, result, test.expected)
		}
	}
}

func TestAddStudent(t *testing.T) {
	student := models.Student{
		ID:        "12345",
		FullName:  "John Doe",
		BirthDate: "01/01/2000",
		Gender:    "Male",
		Faculty:   "Business English",
		Course:    "3",
		Program:   "Computer Science",
		Address:   "123 Main St",
		Email:     "test@go.test",
		Phone:     "1234567890",
		Status:    "Studying",
	}

	input := strings.Join([]string{
		student.ID,
		student.FullName,
		student.BirthDate,
		student.Gender,
		student.Faculty,
		student.Course,
		student.Program,
		student.Address,
		student.Email,
		student.Phone,
		student.Status,
	}, "\n") + "\n"

	var mockInput = strings.NewReader(input)
	var mockOutput bytes.Buffer

	services.AddStudent(mockInput, &mockOutput)

	expectedOutput := "Student added successfully!"
	if !strings.Contains(mockOutput.String(), expectedOutput) {
		t.Errorf("Expected output to contain '%s', but got: %s", expectedOutput, mockOutput.String())
	}
}

func TestDeleteStudent(t *testing.T) {
	services.Students = []models.Student{
		{ID: "123"},
	}

	tests := []struct {
		id       string
		expected string
	}{
		{"123", "Student deleted successfully!"},
		{"456", "Student ID not found."},
	}

	for _, test := range tests {
		input := test.id + "\n"
		var mockInput = strings.NewReader(input)
		var mockOutput bytes.Buffer

		services.DeleteStudent(mockInput, &mockOutput)

		if !strings.Contains(mockOutput.String(), test.expected) {
			t.Errorf("Expected output to contain '%s', but got: %s", test.expected, mockOutput.String())
		}
	}
}

func TestUpdateStudent(t *testing.T) {
	services.Students = []models.Student{
		{
			ID:       "001",
			FullName: "John Doe",
			Email:    "old@example.com",
			Phone:    "0123456789",
			Status:   "Studying",
		},
	}

	tests := []struct {
		id       string
		email    string
		phone    string
		status   string
		expected string
	}{
		{"001", "new@example.com", "0987654321", "Graduated", "Student updated successfully!"},
		{"002", "new@example.com", "0987654321", "Graduated", "Student ID not found."},
	}

	for _, test := range tests {
		input := strings.Join([]string{test.id, test.email, test.phone, test.status}, "\n") + "\n"
		var mockInput = strings.NewReader(input)
		var mockOutput bytes.Buffer

		services.UpdateStudent(mockInput, &mockOutput)

		if !strings.Contains(mockOutput.String(), test.expected) {
			t.Errorf("Expected output to contain '%s', but got: %s", test.expected, mockOutput.String())
		}
	}
}

func TestSearchStudent(t *testing.T) {
	services.Students = []models.Student{
		{
			ID:       "001",
			FullName: "John Doe",
			Email:    "john@example.com",
			Phone:    "0123456789",
			Status:   "Studying",
		},
		{
			ID:       "002",
			FullName: "Jane Doe",
			Email:    "jane@example.com",
			Phone:    "0987654321",
			Status:   "Graduated",
		},
	}

	tests := []struct {
		keyword  string
		expected string
	}{
		{"001", "ID: 001, Name: John Doe, Email: john@example.com, Phone: 0123456789, Status: Studying"},
		{"Jane", "ID: 002, Name: Jane Doe, Email: jane@example.com, Phone: 0987654321, Status: Graduated"},
		{"003", "No student found."},
	}

	for _, test := range tests {
		input := test.keyword + "\n"
		var mockInput = strings.NewReader(input)
		var mockOutput bytes.Buffer

		services.SearchStudent(mockInput, &mockOutput)

		if !strings.Contains(mockOutput.String(), test.expected) {
			t.Errorf("Expected output to contain '%s', but got: %s", test.expected, mockOutput.String())
		}
	}
}
