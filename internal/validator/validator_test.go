package validator

import (
	"os-schema-check/internal/schema"
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		schema schema.Schema
		doc    map[string]any
		want   Result
	}{
		{
			name:   "empty schema, empty document",
			schema: schema.Schema{},
			doc:    map[string]any{},
			want:   Result{Valid: true},
		},
		{
			name: "document contains field not declared in schema",
			schema: schema.Schema{
				Mappings: schema.Mappings{
					Properties: map[string]schema.Property{
						"title": {Type: "text"},
					},
				},
			},
			doc: map[string]any{
				"title": "Seven",
				"extra": "unexpected",
			},
			want: Result{
				Valid:  false,
				Issues: []Issue{{Field: "extra", Message: "unexpected field"}},
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Validate(tc.schema, tc.doc)
			if got.Valid != tc.want.Valid {
				t.Fatalf("Valid = %v, want %v", got.Valid, tc.want.Valid)
			}
			if len(got.Issues) != len(tc.want.Issues) {
				t.Fatalf("Issues length = %d, want %d", len(got.Issues), len(tc.want.Issues))
			}
			if !reflect.DeepEqual(got.Issues, tc.want.Issues) {
				t.Fatalf("Issues = %#v, want %#v", got.Issues, tc.want.Issues)
			}
		})
	}
}
