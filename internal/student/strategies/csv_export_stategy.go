package strategies

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
)

type CSVExportStrategy struct{}

func (s *CSVExportStrategy) Export(ctx context.Context, students []models.Student, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Ghi tiêu đề (header)
	writer.Write([]string{"ID", "StudentID", "FullName", "BirthDate", "GenderID", "FacultyID", "CourseID", "ProgramID", "Address", "Email", "Phone", "StatusID"})

	// Ghi dữ liệu
	for _, student := range students {
		writer.Write([]string{
			strconv.Itoa(int(student.ID)),
			student.StudentID,
			student.FullName,
			student.BirthDate.Format("2006-01-02"),
			strconv.Itoa(student.GenderID),
			strconv.Itoa(student.FacultyID),
			strconv.Itoa(student.CourseID),
			strconv.Itoa(student.ProgramID),
			student.Address,
			student.Email,
			student.Phone,
			strconv.Itoa(student.StatusID),
		})
	}

	return nil
}
