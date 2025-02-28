package http

import (
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
)

func TestNewFileProcessingHandlers(t *testing.T) {
	type args struct {
		fu fileprocessor.IFileProcessorUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want fileprocessor.IFileProcessorHandlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileProcessingHandlers(tt.args.fu, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileProcessingHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileProcessingHandlers_ImportFile(t *testing.T) {
	type fields struct {
		fu fileprocessor.IFileProcessorUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fh := &fileProcessingHandlers{
				fu: tt.fields.fu,
				lg: tt.fields.lg,
			}
			if got := fh.ImportFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileProcessingHandlers.ImportFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileProcessingHandlers_ExportFile(t *testing.T) {
	type fields struct {
		fu fileprocessor.IFileProcessorUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fh := &fileProcessingHandlers{
				fu: tt.fields.fu,
				lg: tt.fields.lg,
			}
			if got := fh.ExportFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileProcessingHandlers.ExportFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
