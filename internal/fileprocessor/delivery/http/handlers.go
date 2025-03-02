package http

import (
	"io"
	"net/http"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/csc13010-student-management/internal/fileprocessor/processor"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type fileProcessingHandlers struct {
	fu fileprocessor.IFileProcessorUsecase
	lg *logger.LoggerZap
}

func NewFileProcessingHandlers(
	fu fileprocessor.IFileProcessorUsecase,
	lg *logger.LoggerZap,
) fileprocessor.IFileProcessorHandlers {
	return &fileProcessingHandlers{
		fu: fu,
		lg: lg,
	}
}

func (fh *fileProcessingHandlers) ImportFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "fileprocessor.ImportFile")
		defer span.Finish()

		module := c.Query("module")
		format := c.Query("format")

		file, err := c.FormFile("file")
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		fileData, err := file.Open()
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}
		defer fileData.Close()

		fileBytes, err := io.ReadAll(fileData)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		err = fh.fu.ImportFile(ctx, module, format, fileBytes)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, nil)
	}
}

func (fh *fileProcessingHandlers) ExportFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "fileprocessor.ExportFile")
		defer span.Finish()

		module := c.Query("module")
		format := c.Query("format")

		data, err := fh.fu.ExportFile(ctx, module, format)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		contentType, ext := processor.GetFileContentType(format)
		filename := module + ext

		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Data(http.StatusOK, contentType, data)
	}
}

// func (s *studentHandlers) ExportStudentCertificate() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		studentID := c.Param("student_id")
// 		format := c.Query("format") // "pdf" hoặc "docx"

// 		student, err := s.su.GetFullInfoStudentByStudentID(c, studentID)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
// 			return
// 		}

// 		var filePath string
// 		if format == "pdf" {
// 			filePath = generatePDF(student)
// 		} else if format == "docx" {
// 			filePath = generateDOCX(student)
// 		} else {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
// 			return
// 		}

// 		c.File(filePath)
// 	}
// }

// func generatePDF(student *dtos.StudentDTO) string {
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()
// 	pdf.SetFont("Arial", "B", 14)

// 	// Tiêu đề trường
// 	pdf.Cell(190, 10, "TRUONG DAI HOC [TEN TRUONG]")
// 	pdf.Ln(6)
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.Cell(190, 10, "PHONG CONG TAC SINH VIEN")
// 	pdf.Ln(10)

// 	// Tiêu đề chính
// 	pdf.SetFont("Arial", "B", 16)
// 	pdf.Cell(190, 10, "GIAY XAC NHAN TINH TRANG SINH VIEN")
// 	pdf.Ln(12)

// 	// Thông tin sinh viên
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.Cell(190, 10, fmt.Sprintf("Ho ten: %s", student.FullName))
// 	pdf.Ln(6)
// 	pdf.Cell(190, 10, fmt.Sprintf("MSSV: %s", student.StudentID))
// 	pdf.Ln(6)
// 	pdf.Cell(190, 10, fmt.Sprintf("Ngay sinh: %s", student.BirthDate))
// 	pdf.Ln(6)
// 	pdf.Cell(190, 10, fmt.Sprintf("Gioi tinh: %s", student.Gender))
// 	pdf.Ln(6)
// 	pdf.Cell(190, 10, fmt.Sprintf("Khoa: %s", student.Faculty))
// 	pdf.Ln(6)
// 	pdf.Cell(190, 10, fmt.Sprintf("Chuong trinh dao tao: %s", student.Program))
// 	pdf.Ln(6)

// 	// Tình trạng sinh viên
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.Cell(190, 10, "Tinh trang sinh vien:")
// 	pdf.Ln(6)
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.Cell(190, 10, strconv.Itoa(student.Status))
// 	pdf.Ln(6)

// 	// Mục đích xác nhận
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.Cell(190, 10, "Muc dich xac nhan:")
// 	pdf.Ln(6)
// 	pdf.SetFont("Arial", "", 12)
// 	// pdf.MultiCell(190, 6, student.Purpose, "", "L", false)
// 	pdf.Ln(6)

// 	// Thời gian hiệu lực
// 	// pdf.Cell(190, 10, fmt.Sprintf("Giay xac nhan co hieu luc den ngay: %s", student.ValidUntil))
// 	pdf.Ln(12)

// 	// Ký tên xác nhận
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.Cell(190, 10, "Xac nhan cua Truong Dai Hoc Khoa Hoc Tu Nhien")
// 	pdf.Ln(12)
// 	pdf.Cell(190, 10, "Ngay cap: "+time.Now().Format("02/01/2006"))
// 	pdf.Ln(18)
// 	pdf.Cell(190, 10, "Truong Phong Dao Tao")
// 	pdf.Ln(6)
// 	pdf.Cell(190, 10, "(Ky, ghi ro ho ten, dong dau)")

// 	// Xuất file
// 	filePath := fmt.Sprintf("exports/student_%s.pdf", student.StudentID)
// 	pdf.OutputFileAndClose(filePath)
// 	return filePath
// }

// func generateDOCX(student *dtos.StudentDTO) string {
// 	doc := document.New()

// 	p := doc.AddParagraph()
// 	r := p.AddRun()
// 	r.AddText("TRUONG DAI HOC [TEN TRUONG]")
// 	r.Properties().SetBold(true)
// 	r.Properties().SetSize(28)

// 	p = doc.AddParagraph()
// 	r = p.AddRun()
// 	r.AddText("PHONG CONG TAC SINH VIEN")
// 	r.Properties().SetSize(12)

// 	p = doc.AddParagraph()
// 	r = p.AddRun()
// 	r.AddText("GIAY XAC NHAN TINH TRANG SINH VIEN")
// 	r.Properties().SetBold(true)
// 	r.Properties().SetSize(28)

// 	p = doc.AddParagraph()
// 	r = p.AddRun()
// 	r.AddText(fmt.Sprintf("Ho ten: %s", student.FullName))

// 	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("MSSV: %s", student.StudentID))
// 	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Ngay sinh: %s", student.BirthDate))
// 	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Gioi tinh: %s", student.Gender))
// 	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Khoa: %s", student.Faculty))
// 	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Chuong trinh dao tao: %s", student.Program))

// 	doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Tinh trang sinh vien: %d", student.Status))
// 	// doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Muc dich xac nhan: %s", student.Purpose))
// 	// doc.AddParagraph().AddRun().AddText(fmt.Sprintf("Giay xac nhan co hieu luc den ngay: %s", student.ValidUntil))

// 	filePath := fmt.Sprintf("exports/student_%s.docx", student.StudentID)
// 	doc.SaveToFile(filePath)
// 	return filePath
// }
