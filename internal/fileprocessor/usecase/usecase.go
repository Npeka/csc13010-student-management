package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/csc13010-student-management/internal/fileprocessor/processor"
	"github.com/csc13010-student-management/pkg/logger"
)

type fileProcessingUseCase struct {
	fr fileprocessor.IFileProcessorRepository
	lg *logger.LoggerZap
}

func NewFileProcessorUsecase(
	fr fileprocessor.IFileProcessorRepository,
	lg *logger.LoggerZap,
) fileprocessor.IFileProcessorUsecase {
	return &fileProcessingUseCase{
		fr: fr,
		lg: lg,
	}
}

func (fu *fileProcessingUseCase) ImportFile(ctx context.Context, module string, format string, fileData []byte) error {
	processor, err := processor.NewFileProcessor(format)
	if err != nil {
		return err
	}

	data, err := processor.Import(fileData)
	if err != nil {
		return err
	}

	return fu.fr.SaveImportedData(ctx, module, data)
}

func (fu *fileProcessingUseCase) ExportFile(ctx context.Context, module string, format string) ([]byte, error) {
	data, err := fu.fr.GetExportData(ctx, module)
	if err != nil {
		return nil, err
	}

	processor, err := processor.NewFileProcessor(format)
	if err != nil {
		return nil, err
	}

	return processor.Export(data)
}
