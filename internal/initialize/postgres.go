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
func NewPostgres(cfg config.PostgresConfig) *gorm.DB {
	// Connection string (DSN)
	fmt.Println(cfg)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		cfg.Host, cfg.Username, cfg.Password, cfg.Dbname, cfg.Port,
	)

	// Configure logger for GORM
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // Standard logger
		logger.Config{
			SlowThreshold:             time.Second, // Log queries slower than 1 second
			LogLevel:                  logger.Info, // Log detailed information
			IgnoreRecordNotFoundError: true,        // Ignore Record Not Found error
			Colorful:                  true,        // Colorful output in terminal
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

	log.Println("Connected to PostgreSQL")
	return db
}
