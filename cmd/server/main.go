package main

import (
	"fmt"
	"log"

	"github.com/csc13010-student-management/internal/initialize"
	"github.com/csc13010-student-management/internal/migration"
	"github.com/csc13010-student-management/internal/server"
)

func main() {
	cfg := initialize.LoadConfig()
	fmt.Println(cfg)
	lg := initialize.NewLogger(cfg.Logger)
	pg := initialize.NewPostgres(cfg.Postgres)
	rd := initialize.NewRedis(cfg.Redis)

	migration.Migrate(pg)

	server := server.NewServer(cfg, lg, pg, rd)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
