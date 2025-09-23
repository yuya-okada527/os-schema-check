package fileutil

import "testing"

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
