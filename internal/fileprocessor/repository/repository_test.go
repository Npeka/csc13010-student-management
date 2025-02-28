package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"gorm.io/gorm"
)

func TestNewFileProcessorRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want fileprocessor.IFileProcessorRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileProcessorRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileProcessorRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileProcessingRepository_SaveImportedData(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		module string
		data   []map[string]interface{}
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
			fr := &fileProcessingRepository{
				db: tt.fields.db,
			}
			if err := fr.SaveImportedData(tt.args.ctx, tt.args.module, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("fileProcessingRepository.SaveImportedData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileProcessingRepository_GetExportData(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		module string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &fileProcessingRepository{
				db: tt.fields.db,
			}
			got, err := fr.GetExportData(tt.args.ctx, tt.args.module)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileProcessingRepository.GetExportData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileProcessingRepository.GetExportData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToStruct(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    T
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToStruct(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
