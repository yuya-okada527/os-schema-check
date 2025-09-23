package schema

import (
	"os-schema-check/internal/fileutil"
	"fmt"
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
	return Schema{}, nil
}