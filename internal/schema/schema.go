package schema

import (
	"os-schema-check/internal/fileutil"
	"fmt"
	"os"
	"encoding/json"
)

type Property struct {
    Type string `json:"type"`
}

type Mappings struct {
    Properties map[string]Property `json:"properties"`
}

type Schema struct {
    Mappings Mappings `json:"mappings"`
}

func Load(path string) (Schema, error) {
	if !fileutil.CheckExtension(path, ".json") {
		return Schema{}, fmt.Errorf("file must have .json extension")
	}
	if !fileutil.IsAvailable(path) {
		return Schema{}, fmt.Errorf("file does not exist: %s", path)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Schema{}, fmt.Errorf("failed to read file: %s", path)
	}
	var schema Schema
	if err := json.Unmarshal(content, &schema); err != nil {
		return Schema{}, fmt.Errorf("failed to parse JSON: %v", err)
	}
	return schema, nil
}