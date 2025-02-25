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
		}, "name", func(p models.Program) interface{} { return p.Name })
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "faculties", []models.Faculty{
			{Name: "Law"},
			{Name: "Business English"},
			{Name: "Japanese"},
			{Name: "French"},
		}, "name", func(f models.Faculty) interface{} { return f.Name })
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "courses", []models.Course{
			{Name: "CSC13010 - Software Engineering"},
			{Name: "CSC13011 - Database Management System"},
			{Name: "CSC13012 - Data Structure and Algorithms"},
			{Name: "CSC13013 - Computer Network"},
		}, "name", func(c models.Course) interface{} { return c.Name })
	}()

	go func() {
		defer wg.Done()
		autoMigrateTable(db, "statuses", []models.Status{
			{Name: "Studying"},
			{Name: "Graduated"},
			{Name: "Dropped Out"},
			{Name: "Paused"},
		}, "name", func(s models.Status) interface{} { return s.Name })
	}()

	wg.Wait()

	autoMigrateTable(db, "students", []models.Student{
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
	}, "id = ?", func(s models.Student) interface{} { return s.ID })
}
