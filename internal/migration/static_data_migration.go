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
		"status_transitions": []models.StatusTransition{
			{CurrentStatusID: 1, NewStatusID: 2}, // Studying -> Graduated
			{CurrentStatusID: 1, NewStatusID: 3}, // Studying -> Dropped Out
			{CurrentStatusID: 1, NewStatusID: 4}, // Studying -> Paused
			{CurrentStatusID: 1, NewStatusID: 6}, // Studying -> Suspended
			{CurrentStatusID: 2, NewStatusID: 5}, // Graduated -> Completed, Awaiting Graduation
			{CurrentStatusID: 4, NewStatusID: 1}, // Paused -> Studying
			{CurrentStatusID: 5, NewStatusID: 2}, // Completed, Awaiting Graduation -> Graduated
			{CurrentStatusID: 6, NewStatusID: 1}, // Suspended -> Studying
		},
		"roles": []models.Role{
			{Name: "admin"},
			{Name: "student"},
			{Name: "teacher"},
		},
		"configs": []models.Config{
			{
				ID:            1,
				EmailDomain:   true,
				ValidatePhone: true,
				StatusRules:   true,
				DeleteLimit:   true,
			},
		},
	}

	for table, data := range staticData {
		db.Table(table).Create(data)
	}
}
