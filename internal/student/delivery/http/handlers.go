package http

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		s.lg.Info("GetStudents called")
		students, err := s.su.GetStudents(c)
		if err != nil {
			s.lg.Error("Error getting students", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		s.lg.Info("GetStudents successful")
		c.JSON(http.StatusOK, students)
	}
}

// GetStudentByStudentID implements student.IStudentHandlers.
func (s *studentHandlers) GetStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetStudentByStudentID called")
		student_id := c.Param("student_id")

		student, err := s.su.GetStudentByStudentID(c, student_id)
		if err != nil {
			s.lg.Error("Error getting student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetStudentByStudentID successful")
		c.JSON(http.StatusOK, student)
	}
}

// CreateStudent implements student.IStudentHandlers.
func (s *studentHandlers) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("CreateStudent called")
		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			s.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.su.CreateStudent(c, &student)
		if err != nil {
			s.lg.Error("Error creating student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("CreateStudent successful")
		c.JSON(http.StatusCreated, student)
	}
}

// UpdateStudent implements student.IStudentHandlers.
func (s *studentHandlers) UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("UpdateStudent called")
		id, err := strconv.ParseInt(c.Param("student_id"), 10, 64)
		if err != nil {
			s.lg.Error("Invalid student ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student := models.Student{ID: int(id)}
		if err := c.ShouldBindJSON(&student); err != nil {
			s.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = s.su.UpdateStudent(c, &student)
		if err != nil {
			s.lg.Error("Error updating student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("UpdateStudent successful")
		c.JSON(http.StatusOK, student)
	}
}

// DeleteStudent implements student.IStudentHandlers.
func (s *studentHandlers) DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("DeleteStudent called")
		student_id := c.Param("student_id")

		err := s.su.DeleteStudent(c, student_id)
		if err != nil {
			s.lg.Error("Error deleting student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("DeleteStudent successful")
		c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
	}
}

func (s *studentHandlers) GetOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetOptions called")
		options, err := s.su.GetOptions(c)
		if err != nil {
			s.lg.Error("Error getting options", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetOptions successful")
		c.JSON(http.StatusOK, options)
	}
}

func (s *studentHandlers) ImportStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("ImportStudents called")

		// Nhận file từ frontend
		file, err := c.FormFile("students")
		if err != nil {
			s.lg.Error("Failed to get uploaded file", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
			return
		}

		// Lưu file tạm
		tempFilePath := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
			s.lg.Error("Failed to save uploaded file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Gọi usecase để xử lý file
		err = s.su.ImportStudents(c, tempFilePath)
		if err != nil {
			s.lg.Error("Failed to import students", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("ImportStudents successful")
		c.JSON(http.StatusOK, gin.H{"message": "Students imported successfully"})
	}
}

func (s *studentHandlers) ExportStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("ExportStudents called")

		format := c.Query("format") // "csv" hoặc "json"
		if format != "csv" && format != "json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
			return
		}

		// Gọi Usecase để lấy đường dẫn file export
		filePath, err := s.su.ExportStudents(c, format)
		if err != nil {
			s.lg.Error("Error exporting students", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Mở file export để gửi về client
		file, err := os.Open(filePath)
		if err != nil {
			s.lg.Error("Error opening export file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		// Thiết lập header để trình duyệt tải file
		c.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Transfer-Encoding", "binary")

		// Gửi file về client
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			s.lg.Error("Error sending export file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("ExportStudents successful")
	}
}
