package main

import (
	"log"

	"github.com/csc13010-student-management/internal/initialize"
	"github.com/csc13010-student-management/internal/migration"
	"github.com/csc13010-student-management/internal/server"
)

func main() {
	cfg := initialize.LoadConfig()
	lg := initialize.NewLogger(cfg.Logger)
	pg := initialize.NewPostgres(cfg.Postgres)
	rd := initialize.NewRedis(cfg.Redis)
	ef := initialize.NewCasbinEnforcer(pg)
	initialize.NewKafkaTopics(cfg.Kafka)

	migration.Migrate(pg)

	server := server.NewServer(cfg, lg, pg, rd, ef)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
