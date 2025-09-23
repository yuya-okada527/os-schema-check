package schema

import (
	"os"
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

	unreadable := filepath.Join(dir, "unreadable.json")
	if err := os.WriteFile(unreadable, []byte("{}"), 0o600); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	if err := os.Chmod(unreadable, 0o000); err != nil {
		t.Fatalf("failed to chmod temp file: %v", err)
	}

	if _, err := Load(unreadable); err == nil {
		t.Fatalf("expected error when schema file is not readable")
	}

	unparsable := filepath.Join(dir, "invalid.json")
	if err := os.WriteFile(unparsable, []byte("{invalid json}"), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	if _, err := Load(unparsable); err == nil {
		t.Fatalf("expected error when schema file is not valid JSON")
	}
}
