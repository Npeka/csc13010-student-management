package strategies

import (
	"context"
	"encoding/json"
	"os"

	"github.com/csc13010-student-management/internal/models"
)

type JSONImportStrategy struct{}

func (s *JSONImportStrategy) Import(ctx context.Context, filePath string) ([]models.Student, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var students []models.Student
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&students); err != nil {
		return nil, err
	}

	return students, nil
}
