package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/csc13010-student-management/internal/fileprocessor/processor"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
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
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileprocessor.ImportFile")
	defer span.Finish()

	processor, err := processor.NewFileProcessor(format)
	if err != nil {
		return errors.Wrap(err, "fileprocessor.ImportFile.NewFileProcessor")
	}

	data, err := processor.Import(fileData)
	if err != nil {
		return errors.Wrap(err, "fileprocessor.ImportFile.Import")
	}

	return fu.fr.SaveImportedData(ctx, module, data)
}

func (fu *fileProcessingUseCase) ExportFile(ctx context.Context, module string, format string) ([]byte, error) {
	data, err := fu.fr.GetExportData(ctx, module)
	if err != nil {
		return nil, errors.Wrap(err, "fileprocessor.ExportFile.GetExportData")
	}

	processor, err := processor.NewFileProcessor(format)
	if err != nil {
		return nil, errors.Wrap(err, "fileprocessor.ExportFile.NewFileProcessor")
	}

	return processor.Export(data)
}
