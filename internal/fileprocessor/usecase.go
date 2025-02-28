package fileprocessor

import "context"

type IFileProcessorUsecase interface {
	ImportFile(ctx context.Context, module string, format string, fileData []byte) error
	ExportFile(ctx context.Context, module string, format string) ([]byte, error)
}
