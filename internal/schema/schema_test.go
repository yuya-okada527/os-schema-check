package schema

import (
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "schema.jsonl")
	_, err := Load(path)
	if err == nil {
		t.Fatalf("expected error for non-.json extension")
	}
	path = filepath.Join(dir, "schema.json")
	_, err = Load(path)
	if err == nil {
		t.Fatalf("expected error for non-existing file")
	}
}