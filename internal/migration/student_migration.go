package migration

import (
	"sync"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

func autoMigrateStudent(db *gorm.DB) {
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "genders", []models.Gender{
			{Name: "Male"},
			{Name: "Female"},
			{Name: "Other"},
		})
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "programs", []models.Program{
			{Name: "High Quality"},
			{Name: "Regular"},
			{Name: "Talented Bachelor"},
			{Name: "Advanced Program"},
		})
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "faculties", []models.Faculty{
			{Name: "Law"},
			{Name: "Business English"},
			{Name: "Japanese"},
			{Name: "French"},
		})
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "courses", []models.Course{
			{Name: "CSC13010 - Software Engineering"},
			{Name: "CSC13011 - Database Management System"},
			{Name: "CSC13012 - Data Structure and Algorithms"},
			{Name: "CSC13013 - Computer Network"},
		})
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "statuses", []models.Status{
			{Name: "Studying"},
			{Name: "Graduated"},
			{Name: "Dropped Out"},
			{Name: "Paused"},
			{Name: "Completed, Awaiting Graduation"},
			{Name: "Suspended"},
			{Name: "Other"},
		})
	}()

	wg.Wait()

	autoMigrateTable(db, "students", []models.Student{
		{
			StudentID: "22127180",
			FullName:  "Nguyen Phuc Khang",
			BirthDate: time.Date(2004, 8, 27, 0, 0, 0, 0, time.UTC),
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
			BirthDate: time.Date(2004, 10, 19, 0, 0, 0, 0, time.UTC),
			GenderID:  2, // Female
			FacultyID: 2,
			CourseID:  3,
			ProgramID: 1,
			Address:   "HCM",
			Email:     "huynhyenngoc@student.university.edu.vn", // Sửa email hợp lệ
			Phone:     "0903123456",                             // Số hợp lệ VN
			StatusID:  1,
		},
		{
			StudentID: "22127419",
			FullName:  "Nguyen Minh Toan",
			BirthDate: time.Date(2004, 1, 8, 0, 0, 0, 0, time.UTC),
			GenderID:  1, // Male
			FacultyID: 3,
			CourseID:  2,
			ProgramID: 3,
			Address:   "HCM",
			Email:     "minhtoan@student.university.edu.vn", // Sửa lại email đúng format
			Phone:     "0356123456",                         // Số hợp lệ VN
			StatusID:  3,
		},
	})
}
