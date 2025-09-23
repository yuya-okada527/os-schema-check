package document

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDocumentReader_Next(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.jsonl")

	contents := "{\"name\":\"Alice\"}\n{\"name\":\"Bob\"}\n"
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	reader, err := NewDocumentReader(path)
	if err != nil {
		t.Fatalf("NewReader returned error: %v", err)
	}
	defer reader.Close()

	if !reader.Next() {
		t.Fatalf("expected first Next to succeed: %v", reader.Err())
	}
	if got := reader.Current()["name"]; got != "Alice" {
		t.Fatalf("unexpected first object: %v", reader.Current())
	}

	if !reader.Next() {
		t.Fatalf("expected second Next to succeed: %v", reader.Err())
	}
	if got := reader.Current()["name"]; got != "Bob" {
		t.Fatalf("unexpected second object: %v", reader.Current())
	}

	if reader.Next() {
		t.Fatalf("expected iteration to end")
	}
	if err := reader.Err(); err != nil {
		t.Fatalf("unexpected error at end: %v", err)
	}
}

func TestNewDocumentReader_InvalidExtension(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.json")

	contents := "{\"name\":\"Alice\"}\n"
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	if _, err := NewDocumentReader(path); err == nil {
		t.Fatalf("expected extension validation error")
	}
}

func TestNewDocumentReader_FileUnavailable(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.jsonl")

	if _, err := NewDocumentReader(path); err == nil {
		t.Fatalf("expected file availability error")
	}
}

func TestReader_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.jsonl")

	contents := "{invalid}\n"
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	reader, err := NewDocumentReader(path)
	if err != nil {
		t.Fatalf("NewReader returned error: %v", err)
	}
	defer reader.Close()

	if reader.Next() {
		t.Fatalf("expected Next to fail for invalid JSON")
	}
	if reader.Err() == nil {
		t.Fatalf("expected error to be reported")
	}
}

func TestDocumentReader_SkipsBulkActionLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.jsonl")

	contents := "{\"index\":{\"_index\":\"movies\",\"_id\":\"tt1979320\"}}\n" +
		"{\"title\":\"Seven\"}\n" +
		"{\"index\":{\"_index\":\"movies\",\"_id\":\"tt0114369\"}}\n" +
		"{\"title\":\"Se7en\"}\n"
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	reader, err := NewDocumentReader(path)
	if err != nil {
		t.Fatalf("NewReader returned error: %v", err)
	}
	defer reader.Close()

	if !reader.Next() {
		t.Fatalf("expected first document: %v", reader.Err())
	}
	if got := reader.Current()["title"]; got != "Seven" {
		t.Fatalf("unexpected first document: %v", reader.Current())
	}

	if !reader.Next() {
		t.Fatalf("expected second document: %v", reader.Err())
	}
	if got := reader.Current()["title"]; got != "Se7en" {
		t.Fatalf("unexpected second document: %v", reader.Current())
	}

	if reader.Next() {
		t.Fatalf("expected no more documents")
	}
	if err := reader.Err(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDocumentReader_MissingDocumentAfterAction(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.jsonl")

	contents := "{\"index\":{\"_index\":\"movies\"}}\n"
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	reader, err := NewDocumentReader(path)
	if err != nil {
		t.Fatalf("NewReader returned error: %v", err)
	}
	defer reader.Close()

	if reader.Next() {
		t.Fatalf("expected Next to fail without trailing document")
	}
	if reader.Err() == nil {
		t.Fatalf("expected error to be reported")
	}
}

func TestDocumentReader_DeleteActionDoesNotExpectDocument(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.jsonl")

	contents := "{\"index\":{\"_index\":\"movies\",\"_id\":\"tt1979320\"}}\n" +
		"{\"title\":\"Seven\"}\n" +
		"{\"delete\":{\"_index\":\"movies\",\"_id\":\"tt0114369\"}}\n"
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	reader, err := NewDocumentReader(path)
	if err != nil {
		t.Fatalf("NewReader returned error: %v", err)
	}
	defer reader.Close()

	if !reader.Next() {
		t.Fatalf("expected first document: %v", reader.Err())
	}
	if got := reader.Current()["title"]; got != "Seven" {
		t.Fatalf("unexpected first document: %v", reader.Current())
	}

	if reader.Next() {
		t.Fatalf("expected no more documents")
	}
	if err := reader.Err(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
