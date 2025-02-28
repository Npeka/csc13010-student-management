package processor

import (
	"fmt"

	"github.com/csc13010-student-management/internal/fileprocessor"
)

var fileContentTypes = map[string]struct {
	MimeType  string
	Extension string
}{
	"csv":  {"text/csv", ".csv"},
	"json": {"application/json", ".json"},
	"xlsx": {"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", ".xlsx"},
	"pdf":  {"application/pdf", ".pdf"},
}

func GetFileContentType(format string) (string, string) {
	if val, ok := fileContentTypes[format]; ok {
		return val.MimeType, val.Extension
	}
	return "application/octet-stream", ""
}

func NewFileProcessor(format string) (fileprocessor.IFileProcessor, error) {
	if _, ext := GetFileContentType(format); ext == "" {
		return nil, fmt.Errorf("unsupported file format")
	}

	switch format {
	case "csv":
		return NewCSVProcessor(), nil
	case "json":
		return NewJSONProcessor(), nil
	default:
		return nil, fmt.Errorf("unsupported file format")
	}
}
