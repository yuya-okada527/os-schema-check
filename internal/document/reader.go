package document

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// DocumentReader provides iterator-style access to newline-delimited JSON documents.
type DocumentReader struct {
	file    *os.File
	scanner *bufio.Scanner
	current map[string]any
	err     error
}

// NewDocumentReader opens the given path and prepares a DocumentReader.
func NewDocumentReader(path string) (*DocumentReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open data file: %w", err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	return &DocumentReader{
		file:    f,
		scanner: scanner,
	}, nil
}

// Next advances the reader to the next JSON object.
// It returns true when a new object is available.
func (r *DocumentReader) Next() bool {
	if r.err != nil {
		return false
	}

	expectDocument := false

	for {
		if !r.scanner.Scan() {
			if err := r.scanner.Err(); err != nil {
				r.err = fmt.Errorf("scan data line: %w", err)
			} else if expectDocument {
				r.err = fmt.Errorf("incomplete bulk action: expected document line")
			}
			return false
		}

		var item map[string]any
		if err := json.Unmarshal(r.scanner.Bytes(), &item); err != nil {
			r.err = fmt.Errorf("parse data line: %w", err)
			return false
		}

		if isBulkActionLine(item) {
			expectDocument = true
			continue
		}

		expectDocument = false
		r.current = item
		return true
	}
}

// Current returns the most recently decoded object.
func (r *DocumentReader) Current() map[string]any {
	return r.current
}

// Err returns the first error encountered during iteration.
func (r *DocumentReader) Err() error {
	return r.err
}

// Close releases the underlying file handle.
func (r *DocumentReader) Close() error {
	if r.file == nil {
		return nil
	}
	err := r.file.Close()
	r.file = nil
	return err
}

func isBulkActionLine(item map[string]any) bool {
	if len(item) != 1 {
		return false
	}

	for key, value := range item {
		switch key {
		case "index", "create", "update", "delete":
			if _, ok := value.(map[string]any); ok {
				return true
			}
		}
	}

	return false
}
