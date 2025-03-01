package migration

import (
	"log"

	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

var allModels = []interface{}{
	// student management models
	&models.User{},
	&models.Gender{},
	&models.Program{},
	&models.Faculty{},
	&models.Course{},
	&models.Status{},
	&models.Student{},

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

func autoMigrateTable[T any](db *gorm.DB, tableName string, data []T) {
	db.Table(tableName).Create(&data)
}
