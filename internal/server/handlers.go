package server

import (
	alHttp "github.com/csc13010-student-management/internal/auditlog/delivery/http"
	alRepository "github.com/csc13010-student-management/internal/auditlog/repository"
	alUsecase "github.com/csc13010-student-management/internal/auditlog/usecase"
	ftHttp "github.com/csc13010-student-management/internal/faculty/delivery/http"
	ftRepository "github.com/csc13010-student-management/internal/faculty/repository"
	ftUsecase "github.com/csc13010-student-management/internal/faculty/usecase"
	fpHttp "github.com/csc13010-student-management/internal/fileprocessor/delivery/http"
	fpRepository "github.com/csc13010-student-management/internal/fileprocessor/repository"
	fpUsecase "github.com/csc13010-student-management/internal/fileprocessor/usecase"
	pgHttp "github.com/csc13010-student-management/internal/program/delivery/http"
	pgRepository "github.com/csc13010-student-management/internal/program/repository"
	pgUsecase "github.com/csc13010-student-management/internal/program/usecase"
	stuHttp "github.com/csc13010-student-management/internal/status/delivery/http"
	stuRepository "github.com/csc13010-student-management/internal/status/repository"
	stuUsecase "github.com/csc13010-student-management/internal/status/usecase"
	stdHttp "github.com/csc13010-student-management/internal/student/delivery/http"
	stdRepository "github.com/csc13010-student-management/internal/student/repository"
	stdUsecase "github.com/csc13010-student-management/internal/student/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(r *gin.Engine) error {
	// repository
	stRepo := stdRepository.NewStudentRepository(s.pg)
	stuRepo := stuRepository.NewStatusRepository(s.pg)
	pgRepo := pgRepository.NewProgramRepository(s.pg)
	ftRepo := ftRepository.NewFacultyRepository(s.pg)
	alRepo := alRepository.NewAuditLogRepository(s.pg)
	fpRepo := fpRepository.NewFileProcessorRepository(s.pg)

	// usecase
	stUc := stdUsecase.NewStudentUsecase(stRepo, s.lg)
	stuUc := stuUsecase.NewStatusUsecase(stuRepo, s.lg)
	pgUc := pgUsecase.NewProgramUsecase(pgRepo, s.lg)
	ftUc := ftUsecase.NewFacultyUsecase(ftRepo, s.lg)
	alUc := alUsecase.NewAuditLogUsecase(alRepo, s.lg)
	fpUc := fpUsecase.NewFileProcessorUsecase(fpRepo, s.lg)

	// handler
	stHandler := stdHttp.NewStudentHandlers(stUc, s.lg)
	stuHandler := stuHttp.NewStatusHandlers(stuUc, s.lg)
	pgHandler := pgHttp.NewProgramHandlers(pgUc, s.lg)
	ftHandler := ftHttp.NewFacultyHandlers(ftUc, s.lg)
	alHandler := alHttp.NewAuditLogHandlers(alUc, s.lg)
	fpHandler := fpHttp.NewFileProcessingHandlers(fpUc, s.lg)

	// router group
	v1 := r.Group("/api/v1")
	stdGroup := v1.Group("/students")
	stuGroup := v1.Group("/statuses")
	pgGroup := v1.Group("/programs")
	ftGroup := v1.Group("/faculties")
	alGroup := v1.Group("/auditlogs")
	fpGroup := v1.Group("/fileprocessor")

	// router
	stdHttp.MapStudentHandlers(stdGroup, stHandler)
	stuHttp.MapStatusHandlers(stuGroup, stuHandler)
	pgHttp.MapProgramHandlers(pgGroup, pgHandler)
	ftHttp.MapFacultyHandlers(ftGroup, ftHandler)
	alHttp.MapAuditLogHandlers(alGroup, alHandler)
	fpHttp.MapfileProcessingHandlers(fpGroup, fpHandler)

	return nil
}
