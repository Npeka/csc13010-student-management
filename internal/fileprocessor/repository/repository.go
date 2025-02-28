package repository

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

type fileProcessingRepository struct {
	db *gorm.DB
}

func NewFileProcessorRepository(
	db *gorm.DB,
) fileprocessor.IFileProcessorRepository {
	return &fileProcessingRepository{db: db}
}

func (fr *fileProcessingRepository) SaveImportedData(ctx context.Context, module string, data []map[string]interface{}) error {
	var students []models.Student
	for _, row := range data {
		student, err := ConvertToStruct[models.Student](row)
		if err != nil {
			continue
		}
		students = append(students, student)
	}

	if err := fr.db.WithContext(ctx).Table(module).Create(students).Error; err != nil {
		return fmt.Errorf("failed to save imported data: %w", err)
	}
	return nil
}

func (fr *fileProcessingRepository) GetExportData(ctx context.Context, module string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	if err := fr.db.WithContext(ctx).Table(module).Find(&results).Error; err != nil {
		return nil, err
	}
	for i := range results {
		delete(results[i], "id")
	}
	return results, nil
}

func ConvertToStruct[T any](data map[string]interface{}) (T, error) {
	var result T
	elem := reflect.ValueOf(&result).Elem()
	typ := elem.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := elem.Field(i)
		structField := typ.Field(i)

		jsonTag := structField.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = structField.Name
		}

		value, exists := data[jsonTag]
		if !exists {
			continue
		}

		strValue, ok := value.(string)
		if !ok {
			continue
		}

		if field.Kind() == reflect.String {
			field.SetString(strValue)
			continue
		}

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			num, err := strconv.Atoi(strValue)
			if err == nil {
				field.SetInt(int64(num))
			}

		case reflect.Float32, reflect.Float64:
			num, err := strconv.ParseFloat(strValue, 64)
			if err == nil {
				field.SetFloat(num)
			}

		case reflect.Bool:
			boolVal, err := strconv.ParseBool(strValue)
			if err == nil {
				field.SetBool(boolVal)
			}

		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				formats := []string{"2006-01-02", "02-01-2006", "2006/01/02"}
				for _, format := range formats {
					if parsedTime, err := time.Parse(format, strValue); err == nil {
						field.Set(reflect.ValueOf(parsedTime))
						break
					}
				}
			}
		}
	}

	return result, nil
}
