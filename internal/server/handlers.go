package server

import (
	alHttp "github.com/csc13010-student-management/internal/auditlog/delivery/http"
	alRepository "github.com/csc13010-student-management/internal/auditlog/repository"
	alUsecase "github.com/csc13010-student-management/internal/auditlog/usecase"
	ftHttp "github.com/csc13010-student-management/internal/faculty/delivery/http"
	ftRepository "github.com/csc13010-student-management/internal/faculty/repository"
	ftUsecase "github.com/csc13010-student-management/internal/faculty/usecase"
	pgHttp "github.com/csc13010-student-management/internal/program/delivery/http"
	pgRepository "github.com/csc13010-student-management/internal/program/repository"
	pgUsecase "github.com/csc13010-student-management/internal/program/usecase"
	stHttp "github.com/csc13010-student-management/internal/student/delivery/http"
	stRepository "github.com/csc13010-student-management/internal/student/repository"
	stUsecase "github.com/csc13010-student-management/internal/student/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(r *gin.Engine) error {
	// repository
	stRepo := stRepository.NewStudentRepository(s.pg)
	alRepo := alRepository.NewAuditLogRepository(s.pg)
	pgRepo := pgRepository.NewProgramRepository(s.pg)
	ftRepo := ftRepository.NewFacultyRepository(s.pg)

	// usecase
	stUc := stUsecase.NewStudentUsecase(stRepo, s.lg)
	alUc := alUsecase.NewAuditLogUsecase(alRepo, s.lg)
	pgUc := pgUsecase.NewProgramUsecase(pgRepo, s.lg)
	ftUc := ftUsecase.NewFacultyUsecase(ftRepo, s.lg)

	// handler
	stHandler := stHttp.NewStudentHandlers(stUc, s.lg)
	alHandler := alHttp.NewAuditLogHandlers(alUc, s.lg)
	pgHandler := pgHttp.NewProgramHandlers(pgUc, s.lg)
	ftHandler := ftHttp.NewFacultyHandlers(ftUc, s.lg)

	// router group
	v1 := r.Group("/api/v1")
	stGroup := v1.Group("/students")
	alGroup := v1.Group("/auditlogs")
	pgGroup := v1.Group("/programs")
	ftGroup := v1.Group("/faculties")

	// router
	stHttp.MapStudentHandlers(stGroup, stHandler)
	alHttp.MapAuditLogHandlers(alGroup, alHandler)
	pgHttp.MapProgramHandlers(pgGroup, pgHandler)
	ftHttp.MapFacultyHandlers(ftGroup, ftHandler)

	return nil
}
