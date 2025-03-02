package migration

import (
	"sync"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/utils/crypto"
	"gorm.io/gorm"
)

func autoMigrateStudent(db *gorm.DB) {
	var wg sync.WaitGroup

	go func() {

		defer wg.Done()
		autoMigrateTable(db, "genders", []models.Gender{
			{Name: "Male"},
			{Name: "Female"},
			{Name: "Other"},
		})
		wg.Add(1)
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "programs", []models.Program{
			{Name: "High Quality"},
			{Name: "Regular"},
			{Name: "Talented Bachelor"},
			{Name: "Advanced Program"},
		})
		wg.Add(1)
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "faculties", []models.Faculty{
			{Name: "Law"},
			{Name: "Business English"},
			{Name: "Japanese"},
			{Name: "French"},
		})
		wg.Add(1)
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "courses", []models.Course{
			{Name: "CSC13010 - Software Engineering"},
			{Name: "CSC13011 - Database Management System"},
			{Name: "CSC13012 - Data Structure and Algorithms"},
			{Name: "CSC13013 - Computer Network"},
		})
		wg.Add(1)
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
		wg.Add(1)
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "users", []models.User{
			{
				Username: "22127180",
				Password: crypto.GetHash("22127180"),
			},
			{
				Username: "22127108",
				Password: crypto.GetHash("22127108"),
			},
			{
				Username: "22127419",
				Password: crypto.GetHash("22127419"),
			},
		})
		wg.Add(1)
	}()

	wg.Wait()

	autoMigrateTable(db, "students", []models.Student{
		{
			UserID:    func(u uint) *uint { return &u }(1),
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
			UserID:    func(u uint) *uint { return &u }(2),
			StudentID: "22127108",
			FullName:  "Huynh Yen Ngoc",
			BirthDate: "2004-11-08",
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
			UserID:    func(u uint) *uint { return &u }(3),
			StudentID: "22127419",
			FullName:  "Nguyen Minh Toan",
			BirthDate: "2004-04-19",
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
