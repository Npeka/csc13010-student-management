package server

import (
	stHttp "github.com/csc13010-student-management/internal/student/delivery/http"
	"github.com/csc13010-student-management/internal/student/repository"
	"github.com/csc13010-student-management/internal/student/usecase"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(r *gin.Engine) error {
	// repository
	stRepo := repository.NewStudentRepository(s.pg)

	// usecase
	stUsecase := usecase.NewStudentUsecase(stRepo, s.lg)

	// handler
	stHandler := stHttp.NewStudentHandler(stUsecase, s.lg)

	// router group
	v1 := r.Group("/api/v1")
	stGroup := v1.Group("/students")

	// router
	stHttp.MapStudentHandlers(stGroup, stHandler)

	return nil
}
