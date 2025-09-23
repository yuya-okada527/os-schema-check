package schema

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	t.Run("rejects non json extension", func(t *testing.T) {
		path := filepath.Join(dir, "schema.jsonl")
		if _, err := Load(path); err == nil {
			t.Fatalf("expected error for non-.json extension")
		}
	})

	t.Run("errors when schema file missing", func(t *testing.T) {
		path := filepath.Join(dir, "missing.json")
		if _, err := Load(path); err == nil {
			t.Fatalf("expected error for non-existing file")
		}
	})

	t.Run("errors when schema file unreadable", func(t *testing.T) {
		unreadable := filepath.Join(dir, "unreadable.json")
		if err := os.WriteFile(unreadable, []byte("{}"), 0o600); err != nil {
			t.Fatalf("failed to write temp file: %v", err)
		}
		if err := os.Chmod(unreadable, 0o000); err != nil {
			t.Fatalf("failed to chmod temp file: %v", err)
		}
		defer os.Chmod(unreadable, 0o600)

		if _, err := Load(unreadable); err == nil {
			t.Fatalf("expected error when schema file is not readable")
		}
	})

	t.Run("errors when schema JSON invalid", func(t *testing.T) {
		invalid := filepath.Join(dir, "invalid.json")
		if err := os.WriteFile(invalid, []byte("{invalid json}"), 0o644); err != nil {
			t.Fatalf("failed to write temp file: %v", err)
		}
		if _, err := Load(invalid); err == nil {
			t.Fatalf("expected error when schema file is not valid JSON")
		}
	})
}
