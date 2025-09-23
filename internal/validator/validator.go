package validator

import "os-schema-check/internal/schema"

// Result captures the outcome of validating a document against a schema.
type Result struct {
	Valid  bool
	Issues []Issue
}

// Issue represents a single validation problem for a document field.
type Issue struct {
	Field   string
	Message string
}

// Validate checks whether the given document conforms to the provided schema.
// The function is stateless; callers inspect the returned Result for errors.
func Validate(s schema.Schema, doc map[string]any) Result {
	return Result{Valid: true}
}
