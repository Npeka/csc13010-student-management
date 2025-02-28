package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/csc13010-student-management/pkg/logger"
)

func TestNewFileProcessorUsecase(t *testing.T) {
	type args struct {
		fr fileprocessor.IFileProcessorRepository
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want fileprocessor.IFileProcessorUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileProcessorUsecase(tt.args.fr, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileProcessorUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileProcessingUseCase_ImportFile(t *testing.T) {
	type fields struct {
		fr fileprocessor.IFileProcessorRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx      context.Context
		module   string
		format   string
		fileData []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &fileProcessingUseCase{
				fr: tt.fields.fr,
				lg: tt.fields.lg,
			}
			if err := fu.ImportFile(tt.args.ctx, tt.args.module, tt.args.format, tt.args.fileData); (err != nil) != tt.wantErr {
				t.Errorf("fileProcessingUseCase.ImportFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileProcessingUseCase_ExportFile(t *testing.T) {
	type fields struct {
		fr fileprocessor.IFileProcessorRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx    context.Context
		module string
		format string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := &fileProcessingUseCase{
				fr: tt.fields.fr,
				lg: tt.fields.lg,
			}
			got, err := fu.ExportFile(tt.args.ctx, tt.args.module, tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileProcessingUseCase.ExportFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileProcessingUseCase.ExportFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
