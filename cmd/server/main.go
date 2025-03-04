package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/csc13010-student-management/internal/initialize"
	"github.com/csc13010-student-management/internal/migration"
	"github.com/csc13010-student-management/internal/server"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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
