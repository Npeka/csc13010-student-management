package fileprocessor

import "context"

type IFileProcessorRepository interface {
	SaveImportedData(ctx context.Context, module string, data []map[string]interface{}) error
	GetExportData(ctx context.Context, module string) ([]map[string]interface{}, error)
}
