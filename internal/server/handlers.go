package server

import (
	alHttp "github.com/csc13010-student-management/internal/auditlog/delivery/http"
	alRepository "github.com/csc13010-student-management/internal/auditlog/repository"
	alUsecase "github.com/csc13010-student-management/internal/auditlog/usecase"
	stHttp "github.com/csc13010-student-management/internal/student/delivery/http"
	stRepository "github.com/csc13010-student-management/internal/student/repository"
	stUsecase "github.com/csc13010-student-management/internal/student/usecase"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(r *gin.Engine) error {
	// repository
	stRepo := stRepository.NewStudentRepository(s.pg)
	alRepo := alRepository.NewAuditLogRepository(s.pg)

	// usecase
	stUsecase := stUsecase.NewStudentUsecase(stRepo, s.lg)
	alUsecase := alUsecase.NewAuditLogUsecase(alRepo, s.lg)

	// handler
	stHandler := stHttp.NewStudentHandlers(stUsecase, s.lg)
	alHandler := alHttp.NewAuditLogHandlers(alUsecase, s.lg)

	// router group
	v1 := r.Group("/api/v1")
	stGroup := v1.Group("/students")
	alGroup := v1.Group("/auditlogs")

	// router
	stHttp.MapStudentHandlers(stGroup, stHandler)
	alHttp.MapAuditLogHandlers(alGroup, alHandler)

	return nil
}
