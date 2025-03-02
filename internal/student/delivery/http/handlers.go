package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type studentHandlers struct {
	su student.IStudentUsecase
	lg *logger.LoggerZap
}

func NewStudentHandlers(
	su student.IStudentUsecase,
	lg *logger.LoggerZap,
) student.IStudentHandlers {
	return &studentHandlers{
		su: su,
		lg: lg,
	}
}

// GetStudents implements student.IStudentHandlers.
func (s *studentHandlers) GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.GetStudents")
		defer span.Finish()

		students, err := s.su.GetStudents(ctx)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, students)
	}
}

// GetStudentByStudentID implements student.IStudentHandlers.
func (s *studentHandlers) GetStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.GetStudentByStudentID")
		defer span.Finish()

		student_id := c.Param("student_id")
		student, err := s.su.GetStudentByStudentID(ctx, student_id)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, student)
	}
}

// GetFullInfoStudentByStudentID implements student.IStudentHandlers.
func (s *studentHandlers) GetFullInfoStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.GetFullInfoStudentByStudentID")
		defer span.Finish()

		student_id := c.Param("student_id")
		student, err := s.su.GetFullInfoStudentByStudentID(ctx, student_id)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, student)
	}
}

// CreateStudent implements student.IStudentHandlers.
func (s *studentHandlers) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.CreateStudent")
		defer span.Finish()

		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := s.su.CreateStudent(ctx, &student)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, student)
	}
}

// UpdateStudent implements student.IStudentHandlers.
func (s *studentHandlers) UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.UpdateStudent")
		defer span.Finish()

		student_id, err := strconv.ParseInt(c.Param("student_id"), 10, 64)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		student := models.Student{StudentID: strconv.FormatInt(student_id, 10)}
		if err := c.ShouldBindJSON(&student); err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err = s.su.UpdateStudent(ctx, &student)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, student)
	}
}

// DeleteStudent implements student.IStudentHandlers.
func (s *studentHandlers) DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.DeleteStudent")
		defer span.Finish()

		s.lg.Info("DeleteStudent called")
		student_id := c.Param("student_id")
		err := s.su.DeleteStudent(ctx, student_id)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, gin.H{"message": "Student deleted successfully"})
	}
}

func (s *studentHandlers) GetOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.GetOptions")
		defer span.Finish()

		options, err := s.su.GetOptions(ctx)
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}
		response.Success(c, response.ErrCodeSuccess, options)
	}
}
