package strategies

import (
	"context"
	"encoding/json"
	"os"

	"github.com/csc13010-student-management/internal/models"
)

type JSONExportStrategy struct{}

func (s *JSONExportStrategy) Export(ctx context.Context, students []models.Student, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(students)
}
