package fileutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckExtension(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		ext      string
		expected bool
	}{
		{
			name:     "matching extension",
			path:     "sample.json",
			ext:      ".json",
			expected: true,
		},
		{
			name:     "different extension",
			path:     "sample.json",
			ext:      ".jsonl",
			expected: false,
		},
		{
			name:     "empty extension",
			path:     "sample.json",
			ext:      "",
			expected: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if actual := CheckExtension(tc.path, tc.ext); actual != tc.expected {
				t.Fatalf("CheckExtension(%q, %q) = %v, expected %v", tc.path, tc.ext, actual, tc.expected)
			}
		})
	}
}

func TestIsAvailable(t *testing.T) {
	dir := t.TempDir()

	existing := filepath.Join(dir, "exists.json")
	if err := os.WriteFile(existing, []byte("{}"), 0o644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	if !IsAvailable(existing) {
		t.Fatalf("expected IsAvailable to return true for existing file")
	}

	missing := filepath.Join(dir, "missing.json")
	if IsAvailable(missing) {
		t.Fatalf("expected IsAvailable to return false for missing file")
	}
}
