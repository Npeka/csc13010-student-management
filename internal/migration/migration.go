package migration

import (
	"log"

	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

var allModels = []interface{}{
	// student management models
	&models.Student{},
	&models.Gender{},
	&models.Program{},
	&models.Faculty{},
	&models.Course{},
	&models.Status{},

	// audit log model
	&models.AuditLog{},
}

func Migrate(db *gorm.DB) {
	log.Println("Start migrating...")

	if err := db.AutoMigrate(allModels...); err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	autoMigrateStudent(db)

	log.Println("Migration completed")
}

// func autoMigrateTable[T any](db *gorm.DB, tableName string, data []T, conditions []string, getValues []func(T) interface{}) {
// 	for _, record := range data {
// 		var count int64
// 		query := ""
// 		args := []interface{}{}

// 		for i, condition := range conditions {
// 			if i > 0 {
// 				query += " AND "
// 			}
// 			query += condition + " = ?"
// 			args = append(args, getValues[i](record))
// 		}

// 		db.Table(tableName).Where(query, args...).Count(&count)

// 		if count == 0 {
// 			db.Create(&record)
// 		}
// 	}
// }

func autoMigrateTable[T any](db *gorm.DB, tableName string, data []T) {
	db.Table(tableName).Create(&data)
}
