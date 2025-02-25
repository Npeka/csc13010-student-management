package migration

import (
	"log"

	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

var allModels = []interface{}{
	// student management models
	&models.Student{},
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

func autoMigrateTable[T any](db *gorm.DB, tableName string, data []T, condition string, getValue func(T) interface{}) {
	for _, record := range data {
		var count int64
		value := getValue(record)
		db.Table(tableName).Where(condition+" = ?", value).Count(&count)

		if count == 0 {
			db.Create(&record)
		}
	}
}
