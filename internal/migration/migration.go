package migration

import (
	"log"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

var allModels = []interface{}{
	&models.Gender{},
	&models.Program{},
	&models.Faculty{},
	&models.Course{},
	&models.Status{},
	&models.StatusTransition{},
	&models.Role{},
	&models.User{},
	&models.Student{},
	&models.Config{},
	&models.AuditLog{},
}

const (
	createExtensionQuery = "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\""
	alterTableQuery      = "ALTER TABLE students REPLICA IDENTITY FULL"
	createIndexQuery     = "CREATE INDEX idx_students_student_id ON students (student_id)"
)

func Migrate(db *gorm.DB) {
	log.Println("Start migrating...")

	executeQuery(db, createExtensionQuery, "creating extension")
	if err := db.AutoMigrate(allModels...); err != nil {
		log.Printf("Error migrating models: %v", err)
	}

	seedStaticData(db)
	executeQuery(db, alterTableQuery, "altering table")
	executeQuery(db, createIndexQuery, "creating index")
	configureDBConnection(db)

	log.Println("Migration completed")
}

func executeQuery(db *gorm.DB, query, action string) {
	if err := db.Exec(query).Error; err != nil {
		log.Printf("Error %s: %v", action, err)
	}
}

func configureDBConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting database connection: %v", err)
		return
	}

	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
