package migration

import (
	"sync"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

func autoMigrateStudent(db *gorm.DB) {
	var wg sync.WaitGroup
	wg.Add(4)

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
		})
	}()

	wg.Wait()

	autoMigrateTable(db, "students", []models.Student{
		{
			StudentID: "22127180",
			FullName:  "Nguyen Phuc Khang",
			BirthDate: time.Date(2004, 8, 27, 0, 0, 0, 0, time.UTC),
			Gender:    models.GenderMap[models.GenderMaleString],
			FacultyID: 1,
			CourseID:  1,
			ProgramID: 1,
			Address:   "HCM",
			Email:     "npkhang287@gmail.com",
			Phone:     "0123456789",
			StatusID:  1,
		},
		{
			StudentID: "22127108",
			FullName:  "Huynh Yen Ngoc",
			BirthDate: time.Date(2004, 10, 19, 0, 0, 0, 0, time.UTC),
			Gender:    models.GenderMap[models.GenderMaleString],
			FacultyID: 1,
			CourseID:  1,
			ProgramID: 1,
			Address:   "HCM",
			Email:     "huynhyenngoc@gmail.com",
			Phone:     "0123456789",
			StatusID:  1,
		},
		{
			StudentID: "22127419",
			FullName:  "Nguyen Minh Toan",
			BirthDate: time.Date(2004, 1, 8, 0, 0, 0, 0, time.UTC),
			Gender:    models.GenderMap[models.GenderMaleString],
			FacultyID: 1,
			CourseID:  1,
			ProgramID: 1,
			Address:   "HCM",
			Email:     "minhtoan@gmail.com",
			Phone:     "0123456789",
			StatusID:  1,
		},
	})
}
