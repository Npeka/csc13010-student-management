package strategies

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/csc13010-student-management/internal/models"
)

type CSVImportStrategy struct{}

func (s *CSVImportStrategy) Import(ctx context.Context, filePath string) ([]models.Student, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var students []models.Student
	for i, row := range rows {
		if i == 0 {
			continue // Bỏ qua header
		}

		birthDate, _ := time.Parse("2006-01-02", row[3])
		genderID, _ := strconv.Atoi(row[4])
		facultyID, _ := strconv.Atoi(row[5])
		courseID, _ := strconv.Atoi(row[6])
		programID, _ := strconv.Atoi(row[7])
		statusID, _ := strconv.Atoi(row[10])

		student := models.Student{
			StudentID: row[1],
			FullName:  row[2],
			BirthDate: birthDate,
			GenderID:  genderID,
			FacultyID: facultyID,
			CourseID:  courseID,
			ProgramID: programID,
			Address:   row[8],
			Email:     row[9],
			Phone:     row[10],
			StatusID:  statusID,
		}

		students = append(students, student)
	}

	return students, nil
}
