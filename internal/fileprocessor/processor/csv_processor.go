package processor

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/csc13010-student-management/internal/fileprocessor"
)

type CSVProcessor struct{}

func NewCSVProcessor() fileprocessor.IFileProcessor {
	return &CSVProcessor{}
}

func (c *CSVProcessor) Import(data []byte) ([]map[string]interface{}, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, value := range row {
			rowMap[headers[i]] = value
		}

		result = append(result, rowMap)
	}

	return result, nil
}

func (c *CSVProcessor) Export(data []map[string]interface{}) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("no data to export")
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	headers := make([]string, 0, len(data[0]))
	for key := range data[0] {
		headers = append(headers, key)
	}

	sort.Strings(headers)

	if err := writer.Write(headers); err != nil {
		return nil, err
	}

	for _, row := range data {
		record := make([]string, len(headers))
		for i, key := range headers {
			record[i] = stringify(row[key])
		}
		if err := writer.Write(record); err != nil {
			return nil, err
		}
	}

	writer.Flush()
	return buf.Bytes(), nil
}

func stringify(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
