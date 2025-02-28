package fileprocessor

type IFileProcessor interface {
	Import(data []byte) ([]map[string]interface{}, error)
	Export(data []map[string]interface{}) ([]byte, error)
}
