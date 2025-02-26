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
		student_id, err := strconv.ParseInt(c.Param("student_id"), 10, 64)
		if err != nil {
			s.lg.Error("Invalid student ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student := models.Student{StudentID: strconv.FormatInt(student_id, 10)}
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

func (s *studentHandlers) GetFaculties() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetFaculties called")
		faculties, err := s.su.GetFaculties(c)
		if err != nil {
			s.lg.Error("Error getting faculties", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetFaculties successful")
		c.JSON(http.StatusOK, faculties)
	}
}

func (s *studentHandlers) GetPrograms() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetPrograms called")
		programs, err := s.su.GetPrograms(c)
		if err != nil {
			s.lg.Error("Error getting programs", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetPrograms successful")
		c.JSON(http.StatusOK, programs)
	}
}

func (s *studentHandlers) GetStatuses() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetFaculties called")
		statuses, err := s.su.GetStatuses(c)
		if err != nil {
			s.lg.Error("Error getting statuses", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetStatuses successful")
		c.JSON(http.StatusOK, statuses)
	}
}

func (s *studentHandlers) CreateFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetFaculties called")
		var faculty models.Faculty
		if err := c.ShouldBindJSON(&faculty); err != nil {
			s.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.su.CreateFaculty(c, &faculty)
		if err != nil {
			s.lg.Error("Error creating faculty", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("CreateFaculty successful")
		c.JSON(http.StatusOK, faculty)
	}
}

func (s *studentHandlers) CreateProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetFaculties called")
		var program models.Program
		if err := c.ShouldBindJSON(&program); err != nil {
			s.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.su.CreateProgram(c, &program)
		if err != nil {
			s.lg.Error("Error creating program", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("CreateProgram successful")
		c.JSON(http.StatusOK, program)
	}
}
func (s *studentHandlers) CreateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetFaculties called")
		var status models.Status
		if err := c.ShouldBindJSON(&status); err != nil {
			s.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.su.CreateStatus(c, &status)
		if err != nil {
			s.lg.Error("Error creating status", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetFaculties successful")
		c.JSON(http.StatusOK, status)
	}
}

func (s *studentHandlers) DeleteFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("DeleteFaculty called")
		faculty_id, err := strconv.Atoi(c.Param("faculty_id"))
		if err != nil {
			s.lg.Error("Invalid faculty ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid faculty ID"})
			return
		}

		err = s.su.DeleteFaculty(c, faculty_id)
		if err != nil {
			s.lg.Error("Error deleting faculty", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("DeleteFaculty successful")
		c.JSON(http.StatusOK, gin.H{"message": "Faculty deleted successfully"})
	}
}

func (s *studentHandlers) DeleteProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("DeleteProgram called")
		program_id, err := strconv.Atoi(c.Param("program_id"))
		if err != nil {
			s.lg.Error("Invalid program ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid program ID"})
			return
		}

		err = s.su.DeleteProgram(c, program_id)
		if err != nil {
			s.lg.Error("Error deleting program", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("DeleteProgram successful")
		c.JSON(http.StatusOK, gin.H{"message": "Program deleted successfully"})
	}
}

func (s *studentHandlers) DeleteStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("DeleteStatus called")
		status_id, err := strconv.Atoi(c.Param("status_id"))
		if err != nil {
			s.lg.Error("Invalid status ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
			return
		}

		err = s.su.DeleteStatus(c, status_id)
		if err != nil {
			s.lg.Error("Error deleting status", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("DeleteStatus successful")
		c.JSON(http.StatusOK, gin.H{"message": "Status deleted successfully"})
	}
}
