package http

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
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

func (s *studentHandlers) GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.GetStudents")
		defer span.Finish()

		students, err := s.su.GetStudents(ctx)
		if err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.OkSuccess, students)
	}
}

func (s *studentHandlers) GetStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.GetStudentByStudentID")
		defer span.Finish()

		student_id := c.Param("student_id")
		student, err := s.su.GetStudentByStudentID(ctx, student_id)
		if err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, student)
	}
}

func (s *studentHandlers) GetFullInfoStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.GetFullInfoStudentByStudentID")
		defer span.Finish()

		student_id := c.Param("student_id")
		student, err := s.su.GetFullInfoStudentByStudentID(ctx, student_id)
		if err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, student)
	}
}

func (s *studentHandlers) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.CreateStudent")
		defer span.Finish()

		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := s.su.CreateStudent(ctx, &student)
		if err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, student)
	}
}

func (s *studentHandlers) UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.UpdateStudent")
		defer span.Finish()

		student_id := c.Param("student_id")
		student := models.Student{StudentID: student_id}
		if err := c.ShouldBindJSON(&student); err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		if err := s.su.UpdateStudent(ctx, &student); err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, nil)
	}
}

func (s *studentHandlers) DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.DeleteStudent")
		defer span.Finish()

		student_id := c.Param("student_id")
		err := s.su.DeleteStudent(ctx, student_id)
		if err != nil {
			logger.LogResponseError(c, s.lg, err)
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
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}
		response.Success(c, response.ErrCodeSuccess, options)
	}
}

func (s *studentHandlers) Export() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "student.Export")
		defer span.Finish()

		studentID := c.Param("student_id")
		ext := c.DefaultQuery("ext", "html")

		student, err := s.su.GetFullInfoStudentByStudentID(ctx, studentID)
		if err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		data := struct {
			University string
			Address    string
			Phone      string
			Email      string
			Student    *dtos.StudentResponseDTO
			ValidUntil string
		}{
			University: "Đại học Công Nghệ",
			Address:    "123 Đường ABC, Quận 1, TP. HCM",
			Phone:      "028-1234-5678",
			Email:      "contact@university.edu.vn",
			Student:    student,
			ValidUntil: time.Now().AddDate(0, 6, 0).Format("02/01/2006"),
		}

		var templateFile string
		var contentType string
		if ext == "md" {
			templateFile = "internal/student/templates/status.md"
			contentType = "text/markdown"
		} else {
			templateFile = "internal/student/templates/status.html"
			contentType = "text/html"
		}

		tmpl, err := template.ParseFiles(templateFile)
		if err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		// Render template
		var rendered bytes.Buffer
		if err := tmpl.Execute(&rendered, data); err != nil {
			logger.LogResponseError(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		// Trả về file tải xuống
		filename := fmt.Sprintf("student_%s.%s", studentID, ext)
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", contentType)
		c.String(http.StatusOK, rendered.String())
	}
}
