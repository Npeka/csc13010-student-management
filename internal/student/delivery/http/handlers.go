package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/unidoc/unioffice/document"
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

const defaultDeleteTimeLimit = 30 * time.Minute

// DeleteStudent implements student.IStudentHandlers.
func (s *studentHandlers) DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("ALLOW_DELETE_ANYTIME") == "false" {
			c.JSON(http.StatusForbidden, gin.H{"error": "DeleteStudent is disabled"})
			return
		}

		s.lg.Info("DeleteStudent called")
		student_id := c.Param("student_id")

		student, err := s.su.GetStudentByStudentID(c, student_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}

		deleteTimeLimit := defaultDeleteTimeLimit
		if envLimit, exists := os.LookupEnv("DELETE_TIME_LIMIT"); exists {
			parsedLimit, err := time.ParseDuration(envLimit)
			if err == nil {
				deleteTimeLimit = parsedLimit
			}
		}

		if time.Since(student.CreatedAt) > deleteTimeLimit {
			c.JSON(http.StatusForbidden, gin.H{"error": "Student cannot be deleted after the allowed time"})
			return
		}

		err = s.su.DeleteStudent(c, student_id)
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



func (s *studentHandlers) GetFullInfoStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetFullInfoStudentByStudentID called")
		student_id := c.Param("student_id")

		student, err := s.su.GetFullInfoStudentByStudentID(c, student_id)
		if err != nil {
			s.lg.Error("Error getting student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetFullInfoStudentByStudentID successful")
		c.JSON(http.StatusOK, student)
	}
}

func (s *studentHandlers) ExportStudentCertificate() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.Param("student_id")
		format := c.Query("format") // "pdf" hoặc "docx"

		student, err := s.su.GetFullInfoStudentByStudentID(c, studentID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}

		var filePath string
		if format == "pdf" {
			filePath = generatePDF(student)
		} else if format == "docx" {
			filePath = generateDOCX(student)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
			return
		}

		c.File(filePath)
	}
}

func generatePDF(student *dtos.StudentDTO) string {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)

	// Tiêu đề trường
	pdf.Cell(190, 10, "TRUONG DAI HOC [TEN TRUONG]")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 10, "PHONG CONG TAC SINH VIEN")
	pdf.Ln(10)

	// Tiêu đề chính
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "GIAY XAC NHAN TINH TRANG SINH VIEN")
	pdf.Ln(12)

	// Thông tin sinh viên
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 10, fmt.Sprintf("Ho ten: %s", student.FullName))
	pdf.Ln(6)
	pdf.Cell(190, 10, fmt.Sprintf("MSSV: %s", student.StudentID))
	pdf.Ln(6)
	pdf.Cell(190, 10, fmt.Sprintf("Ngay sinh: %s", student.BirthDate))
	pdf.Ln(6)
	pdf.Cell(190, 10, fmt.Sprintf("Gioi tinh: %s", student.Gender))
	pdf.Ln(6)
	pdf.Cell(190, 10, fmt.Sprintf("Khoa: %s", student.Faculty))
	pdf.Ln(6)
	pdf.Cell(190, 10, fmt.Sprintf("Chuong trinh dao tao: %s", student.Program))
	pdf.Ln(6)

	// Tình trạng sinh viên
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 10, "Tinh trang sinh vien:")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 10, strconv.Itoa(student.Status))
	pdf.Ln(6)

	// Mục đích xác nhận
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 10, "Muc dich xac nhan:")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 12)
	// pdf.MultiCell(190, 6, student.Purpose, "", "L", false)
	pdf.Ln(6)

	// Thời gian hiệu lực
	// pdf.Cell(190, 10, fmt.Sprintf("Giay xac nhan co hieu luc den ngay: %s", student.ValidUntil))
	pdf.Ln(12)

	// Ký tên xác nhận
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 10, "Xac nhan cua Truong Dai Hoc Khoa Hoc Tu Nhien")
	pdf.Ln(12)
	pdf.Cell(190, 10, "Ngay cap: "+time.Now().Format("02/01/2006"))
	pdf.Ln(18)
	pdf.Cell(190, 10, "Truong Phong Dao Tao")
	pdf.Ln(6)
	pdf.Cell(190, 10, "(Ky, ghi ro ho ten, dong dau)")

	// Xuất file
	filePath := fmt.Sprintf("exports/student_%s.pdf", student.StudentID)
	pdf.OutputFileAndClose(filePath)
	return filePath
}

func generateDOCX(student *dtos.StudentDTO) string {
	doc := document.New()

	p := doc.AddParagraph()
	r := p.AddRun()
	r.AddText("TRUONG DAI HOC [TEN TRUONG]")
	r.Properties().SetBold(true)
	r.Properties().SetSize(28)

	p = doc.AddParagraph()
	r = p.AddRun()
	r.AddText("PHONG CONG TAC SINH VIEN")
	r.Properties().SetSize(12)

	p = doc.AddParagraph()
	r = p.AddRun()
	r.AddText("GIAY XAC NHAN TINH TRANG SINH VIEN")
	r.Properties().SetBold(true)
	r.Properties().SetSize(28)

	p = doc.AddParagraph()
	r = p.AddRun()
	r.AddText(fmt.Sprintf("Ho ten: %s", student.FullName))

	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("MSSV: %s", student.StudentID))
	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Ngay sinh: %s", student.BirthDate))
	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Gioi tinh: %s", student.Gender))
	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Khoa: %s", student.Faculty))
	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Chuong trinh dao tao: %s", student.Program))

	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Tinh trang sinh vien: %d", student.Status))
	// doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Muc dich xac nhan: %s", student.Purpose))
	// doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Giay xac nhan co hieu luc den ngay: %s", student.ValidUntil))

	filePath := fmt.Sprintf("exports/student_%s.docx", student.StudentID)
	doc.SaveToFile(filePath)
	return filePath
}
