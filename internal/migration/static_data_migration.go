package migration

import (
	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

func seedStaticData(db *gorm.DB) {
	staticData := map[string]interface{}{
		"genders": []models.Gender{
			{Name: "Male"},
			{Name: "Female"},
			{Name: "Other"},
		},
		"programs": []models.Program{
			{Name: "High Quality"},
			{Name: "Regular"},
			{Name: "Talented Bachelor"},
			{Name: "Advanced Program"},
		},
		"faculties": []models.Faculty{
			{Name: "Law"},
			{Name: "Business English"},
			{Name: "Japanese"},
			{Name: "French"},
		},
		"courses": []models.Course{
			{Name: "CSC13010 - Software Engineering"},
			{Name: "CSC13011 - Database Management System"},
			{Name: "CSC13012 - Data Structure and Algorithms"},
			{Name: "CSC13013 - Computer Network"},
		},
		"statuses": []models.Status{
			{Name: "Studying"},
			{Name: "Graduated"},
			{Name: "Dropped Out"},
			{Name: "Paused"},
			{Name: "Completed, Awaiting Graduation"},
			{Name: "Suspended"},
			{Name: "Other"},
		},
		"roles": []models.Role{
			{Name: "admin"},
			{Name: "student"},
			{Name: "teacher"},
		},
	}

	for table, data := range staticData {
		db.Table(table).Create(data)
	}
}
