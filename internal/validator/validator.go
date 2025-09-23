package validator

import (
	"os-schema-check/internal/schema"
	"sort"
)

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
	var issues []Issue

	allowed := make(map[string]struct{})
	if s.Mappings.Properties != nil {
		for key := range s.Mappings.Properties {
			allowed[key] = struct{}{}
		}
	}

	keys := make([]string, 0, len(doc))
	for key := range doc {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if _, ok := allowed[key]; ok {
			continue
		}
		issues = append(issues, Issue{
			Field:   key,
			Message: "unexpected field",
		})
	}

	if len(issues) == 0 {
		return Result{Valid: true}
	}

	return Result{Valid: false, Issues: issues}
}
