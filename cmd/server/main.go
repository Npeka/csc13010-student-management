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

	cf := initialize.LoadConfig()
	lg := initialize.NewLogger(cf.Logger)
	pg := initialize.NewPostgres(cf)
	rd := initialize.NewRedis(cf.Redis)
	ef := initialize.NewCasbinEnforcer(pg)
	initialize.NewKafkaTopics(cf.Kafka)

	migration.Migrate(pg)

	server := server.NewServer(cf, lg, pg, rd, ef)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
