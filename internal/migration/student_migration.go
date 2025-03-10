package migration

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"gorm.io/gorm"
)

func SeedStudents(db *gorm.DB, su student.IStudentUsecase) {
	students := []models.Student{
		{
			StudentID: "22127180",
			FullName:  "Nguyen Phuc Khang",
			BirthDate: "2004-10-19",
			GenderID:  1,
			FacultyID: 1,
			CourseID:  1,
			ProgramID: 2,
			Address:   "HCM",
			Email:     "npkhang287@student.university.edu.vn",
			Phone:     "0789123456",
			StatusID:  1,
		},
		{
			StudentID: "22127108",
			FullName:  "Huynh Yen Ngoc",
			BirthDate: "2004-11-08",
			GenderID:  2,
			FacultyID: 2,
			CourseID:  3,
			ProgramID: 1,
			Address:   "HCM",
			Email:     "huynhyenngoc@student.university.edu.vn",
			Phone:     "0903123456",
			StatusID:  1,
		},
		{
			StudentID: "22127419",
			FullName:  "Nguyen Minh Toan",
			BirthDate: "2004-04-19",
			GenderID:  1,
			FacultyID: 3,
			CourseID:  2,
			ProgramID: 3,
			Address:   "HCM",
			Email:     "minhtoan@student.university.edu.vn",
			Phone:     "0356123456",
			StatusID:  3,
		},
	}

	for _, student := range students {
		su.CreateStudent(context.Background(), &student)
	}
}
