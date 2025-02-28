package processor

import (
	"encoding/json"

	"github.com/csc13010-student-management/internal/fileprocessor"
)

type JSONProcessor struct{}

func NewJSONProcessor() fileprocessor.IFileProcessor {
	return &JSONProcessor{}
}

func (j *JSONProcessor) Import(data []byte) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (j *JSONProcessor) Export(data []map[string]interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}
