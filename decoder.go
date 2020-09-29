package urlshort

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Mapper struct {
	filename string
}

type FileDecoder interface {
	Decode(i interface{}) error
}

func newFileDecoder(file *os.File) FileDecoder {
	if strings.Contains(file.Name(), "json") {
		return json.NewDecoder(file)
	}
	return yaml.NewDecoder(file)
}

func (m *Mapper) readFile() (*os.File, error) {
	// Read JSON file mapping
	file, err := os.Open(m.filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (m *Mapper) decode() (map[string]string, error) {
	// Read the file mapping
	file, err := m.readFile()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode
	decoder := newFileDecoder(file)
	var decoded []map[string]string
	err = decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}

	// Process to be a key-value pair / map
	result := make(map[string]string)
	for _, mapping := range decoded {
		result[mapping["path"]] = mapping["url"]
	}

	return result, nil
}
