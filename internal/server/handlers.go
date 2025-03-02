package server

import (
	"fmt"

	alHttp "github.com/csc13010-student-management/internal/auditlog/delivery/http"
	alRepository "github.com/csc13010-student-management/internal/auditlog/repository"
	alUsecase "github.com/csc13010-student-management/internal/auditlog/usecase"
	authHttp "github.com/csc13010-student-management/internal/auth/delivery/http"
	authRepository "github.com/csc13010-student-management/internal/auth/repository"
	authUsecase "github.com/csc13010-student-management/internal/auth/usecase"
	authWorker "github.com/csc13010-student-management/internal/auth/worker"
	ftHttp "github.com/csc13010-student-management/internal/faculty/delivery/http"
	ftRepository "github.com/csc13010-student-management/internal/faculty/repository"
	ftUsecase "github.com/csc13010-student-management/internal/faculty/usecase"
	fpHttp "github.com/csc13010-student-management/internal/fileprocessor/delivery/http"
	fpRepository "github.com/csc13010-student-management/internal/fileprocessor/repository"
	fpUsecase "github.com/csc13010-student-management/internal/fileprocessor/usecase"
	"github.com/csc13010-student-management/internal/initialize"
	"github.com/csc13010-student-management/internal/middleware"
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
	kafkaUtils "github.com/csc13010-student-management/pkg/kafka"
)

func (s *Server) MapHandlers() error {

	// kafka
	kurl := fmt.Sprintf("%s:%d", s.cfg.Kafka.Host, s.cfg.Kafka.Port)
	kwUserCreated := kafkaUtils.NewKafkaWriter(kurl, initialize.KafkaAuthUserCreated)
	kwStudentCreated := kafkaUtils.NewKafkaWriter(kurl, initialize.KafkaStudentCreated)

	// repository
	authRepo := authRepository.NewAuthRepository(s.pg)
	stdRepo := stdRepository.NewStudentRepository(s.pg)
	stuRepo := stuRepository.NewStatusRepository(s.pg)
	pgRepo := pgRepository.NewProgramRepository(s.pg)
	ftRepo := ftRepository.NewFacultyRepository(s.pg)
	alRepo := alRepository.NewAuditLogRepository(s.pg)
	fpRepo := fpRepository.NewFileProcessorRepository(s.pg)

	// usecase
	authUc := authUsecase.NewAuthUsecase(authRepo, s.lg, kwUserCreated, s.e)
	stdUc := stdUsecase.NewStudentUsecase(stdRepo, s.lg, kwStudentCreated, s.e)
	stuUc := stuUsecase.NewStatusUsecase(stuRepo, s.lg)
	pgUc := pgUsecase.NewProgramUsecase(pgRepo, s.lg)
	ftUc := ftUsecase.NewFacultyUsecase(ftRepo, s.lg)
	alUc := alUsecase.NewAuditLogUsecase(alRepo, s.lg)
	fpUc := fpUsecase.NewFileProcessorUsecase(fpRepo, s.lg)

	// handler
	authHandler := authHttp.NewAuthHandlers(authUc, s.lg)
	stHandler := stdHttp.NewStudentHandlers(stdUc, s.lg)
	stuHandler := stuHttp.NewStatusHandlers(stuUc, s.lg)
	pgHandler := pgHttp.NewProgramHandlers(pgUc, s.lg)
	ftHandler := ftHttp.NewFacultyHandlers(ftUc, s.lg)
	alHandler := alHttp.NewAuditLogHandlers(alUc, s.lg)
	fpHandler := fpHttp.NewFileProcessingHandlers(fpUc, s.lg)
	rlHandler := rlHttp.NewRbacHandlers(s.e)

	// worker
	s.w.authWorker = authWorker.NewAuthWorker(authUc)
	s.w.stdWorker = stdWorker.NewStudentWorker(stdUc)

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
	rlHttp.MaprbacHandlers(rlGroup, rlHandler)

	return nil
}
