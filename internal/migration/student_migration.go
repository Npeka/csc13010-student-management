package migration

import (
	"sync"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

func autoMigrateProgram(db *gorm.DB) {
	programs := []models.Program{
		{Name: "High Quality"},
		{Name: "Regular"},
		{Name: "Talented Bachelor"},
		{Name: "Advanced Program"},
	}

	for _, record := range programs {
		var count int64
		db.Table("programs").Where("name = ?", record.Name).Count(&count)
		if count == 0 {
			db.Create(&record)
		}
	}
}

func autoMigrateFaculty(db *gorm.DB) {
	faculties := []models.Faculty{
		{Name: "Law"},
		{Name: "Business English"},
		{Name: "Japanese"},
		{Name: "French"},
	}

	for _, record := range faculties {
		var count int64
		db.Table("faculties").Where("name = ?", record.Name).Count(&count)
		if count == 0 {
			db.Create(&record)
		}
	}
}

func autoMigrateCourse(db *gorm.DB) {
	courses := []models.Course{
		{Name: "CSC13010 - Software Engineering"},
		{Name: "CSC13011 - Database Management System"},
		{Name: "CSC13012 - Data Structure and Algorithms"},
		{Name: "CSC13013 - Computer Network"},
	}

	for _, record := range courses {
		var count int64
		db.Table("courses").Where("name = ?", record.Name).Count(&count)
		if count == 0 {
			db.Create(&record)
		}
	}
}

func autoMigrateStatus(db *gorm.DB) {
	statuses := []models.Status{
		{Name: "Studying"},
		{Name: "Graduated"},
		{Name: "Dropped Out"},
		{Name: "Paused"},
	}

	for _, record := range statuses {
		var count int64
		db.Table("statuses").Where("name = ?", record.Name).Count(&count)
		if count == 0 {
			db.Create(&record)
		}
	}
}

func autoMigrateStudent(db *gorm.DB) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		autoMigrateProgram(db)
	}()

	go func() {
		defer wg.Done()
		autoMigrateFaculty(db)
	}()

	go func() {
		defer wg.Done()
		autoMigrateCourse(db)
	}()

	go func() {
		defer wg.Done()
		autoMigrateStatus(db)
	}()

	wg.Wait()
	seedTable(db, "students", []models.Student{
		{
			ID:        22127180,
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
	}, "id = ?", "22127180")
}
