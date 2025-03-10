package server

import (
	"fmt"

	alHttp "github.com/csc13010-student-management/internal/auditlog/delivery/http"
	alRepository "github.com/csc13010-student-management/internal/auditlog/repository"
	alUsecase "github.com/csc13010-student-management/internal/auditlog/usecase"
	alWorker "github.com/csc13010-student-management/internal/auditlog/worker"
	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/migration"
	"github.com/segmentio/kafka-go"

	authHttp "github.com/csc13010-student-management/internal/auth/delivery/http"
	authRepository "github.com/csc13010-student-management/internal/auth/repository"
	authUsecase "github.com/csc13010-student-management/internal/auth/usecase"
	authWorker "github.com/csc13010-student-management/internal/auth/worker"
	cfHttp "github.com/csc13010-student-management/internal/config/delivery/http"
	cfRepository "github.com/csc13010-student-management/internal/config/repository"
	cfUsecase "github.com/csc13010-student-management/internal/config/usecase"
	ftHttp "github.com/csc13010-student-management/internal/faculty/delivery/http"
	ftRepository "github.com/csc13010-student-management/internal/faculty/repository"
	ftUsecase "github.com/csc13010-student-management/internal/faculty/usecase"
	fpHttp "github.com/csc13010-student-management/internal/fileprocessor/delivery/http"
	fpRepository "github.com/csc13010-student-management/internal/fileprocessor/repository"
	fpUsecase "github.com/csc13010-student-management/internal/fileprocessor/usecase"
	"github.com/csc13010-student-management/internal/middleware"
	notiRepository "github.com/csc13010-student-management/internal/notification/repository"
	notiWorker "github.com/csc13010-student-management/internal/notification/worker"
	pgHttp "github.com/csc13010-student-management/internal/program/delivery/http"
	pgRepository "github.com/csc13010-student-management/internal/program/repository"
	pgUsecase "github.com/csc13010-student-management/internal/program/usecase"
	rlHttp "github.com/csc13010-student-management/internal/rbac/delivery/http"
	stuHttp "github.com/csc13010-student-management/internal/status/delivery/http"
	stuRepository "github.com/csc13010-student-management/internal/status/repository"
	stuUsecase "github.com/csc13010-student-management/internal/status/usecase"
	stdHttp "github.com/csc13010-student-management/internal/student/delivery/http"
	stdRepository "github.com/csc13010-student-management/internal/student/repository"
	stdUsecase "github.com/csc13010-student-management/internal/student/usecase"
	stdWorker "github.com/csc13010-student-management/internal/student/worker"
	kafkas "github.com/csc13010-student-management/pkg/kafka"
)

func (s *Server) MapHandlers() error {
	// kafka
	kurl := fmt.Sprintf("%s:%d", s.cf.Kafka.Host, s.cf.Kafka.Port)
	kwauc := kafkas.NewKafkaWriter(kurl, events.AuthUserCreated)
	kwsstd := map[string]*kafka.Writer{
		events.AuthCreateUser:     kafkas.NewKafkaWriter(kurl, events.AuthCreateUser),
		string(events.NotiCreate): kafkas.NewKafkaWriter(kurl, string(events.NotiCreate)),
	}

	// repository
	authRepo := authRepository.NewAuthRepository(s.pg)
	stdRepo := stdRepository.NewStudentRepository(s.pg)
	stuRepo := stuRepository.NewStatusRepository(s.pg)
	pgRepo := pgRepository.NewProgramRepository(s.pg)
	ftRepo := ftRepository.NewFacultyRepository(s.pg)
	alRepo := alRepository.NewAuditLogRepository(s.pg)
	fpRepo := fpRepository.NewFileProcessorRepository(s.pg)
	notiRepo := notiRepository.NewNotificationRepository(s.pg)
	cfRepo := cfRepository.NewConfigRepository(s.pg)

	// usecase
	authUc := authUsecase.NewAuthUsecase(authRepo, s.lg, kwauc, s.e)
	stdUc := stdUsecase.NewStudentUsecase(stdRepo, s.lg, kwsstd, s.e)
	stuUc := stuUsecase.NewStatusUsecase(stuRepo, s.lg)
	pgUc := pgUsecase.NewProgramUsecase(pgRepo, s.lg)
	ftUc := ftUsecase.NewFacultyUsecase(ftRepo, s.lg)
	alUc := alUsecase.NewAuditLogUsecase(alRepo, s.lg)
	fpUc := fpUsecase.NewFileProcessorUsecase(fpRepo, s.lg)
	cfUc := cfUsecase.NewConfigUsecase(cfRepo, s.lg)

	// handler
	authHandler := authHttp.NewAuthHandlers(authUc, s.lg)
	stHandler := stdHttp.NewStudentHandlers(stdUc, s.lg)
	stuHandler := stuHttp.NewStatusHandlers(stuUc, s.lg)
	pgHandler := pgHttp.NewProgramHandlers(pgUc, s.lg)
	ftHandler := ftHttp.NewFacultyHandlers(ftUc, s.lg)
	alHandler := alHttp.NewAuditLogHandlers(alUc, s.lg)
	fpHandler := fpHttp.NewFileProcessingHandlers(fpUc, s.lg)
	cfHandler := cfHttp.NewConfigHandlers(cfUc, s.lg)
	rlHandler := rlHttp.NewRbacHandlers(s.e)

	// worker
	s.w.authWorker = authWorker.NewAuthWorker(authUc, s.lg)
	s.w.stdWorker = stdWorker.NewStudentWorker(stdUc, s.lg)
	s.w.notiWorker = notiWorker.NewNotificationWorker(notiRepo, s.lg)
	s.w.auditWorker = alWorker.NewAuditLogWorker(alUc, s.lg)

	// migration
	migration.SeedStudents(s.pg, stdUc)

	// router group
	// r.Use(middleware.CasdoorAuthMiddleware())
	// r.Use(middleware.CasbinMiddleware(e))
	v1 := s.r.Group("/api/v1")
	authGroup := v1.Group("/auth")
	stdGroup := v1.Group("/students")
	stuGroup := v1.Group("/statuses")
	pgGroup := v1.Group("/programs")
	ftGroup := v1.Group("/faculties")
	alGroup := v1.Group("/auditlogs")
	fpGroup := v1.Group("/fileprocessor")
	cfGroup := v1.Group("/config")
	rlGroup := v1.Group("/rbac")

	// middleware
	mw := middleware.NewMiddlewareManager(s.e)

	// router
	authHttp.MapAuthHandlers(authGroup, authHandler)
	stdHttp.MapStudentHandlers(stdGroup, stHandler, mw)
	stuHttp.MapStatusHandlers(stuGroup, stuHandler)
	pgHttp.MapProgramHandlers(pgGroup, pgHandler)
	ftHttp.MapFacultyHandlers(ftGroup, ftHandler)
	alHttp.MapAuditLogHandlers(alGroup, alHandler)
	fpHttp.MapfileProcessingHandlers(fpGroup, fpHandler)
	cfHttp.MapconfigHandlers(cfGroup, cfHandler)
	rlHttp.MaprbacHandlers(rlGroup, rlHandler)

	return nil
}
