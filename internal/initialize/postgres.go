package initialize

import (
	"fmt"
	"log"
	"time"

	"github.com/csc13010-student-management/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgres creates a connection to PostgreSQL with GORM
func NewPostgres(cf *config.Config) *gorm.DB {
	pgcf := cf.Postgres

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		pgcf.Host, pgcf.Username, pgcf.Password, pgcf.Dbname, pgcf.Port,
	)

	var logLevel logger.LogLevel
	switch cf.Server.Mode {
	case "dev":
		logLevel = logger.Info
	case "test":
		logLevel = logger.Warn
	case "prod":
		logLevel = logger.Error
	default:
		logLevel = logger.Silent
	}

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  cf.Server.Mode == "dev",
		},
	)

	// Open database connection with GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
		return nil
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Connected to PostgreSQL")
	return db
}
