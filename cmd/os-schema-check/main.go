package main

import (
	"fmt"
	"os"
	"os-schema-check/internal/document"
	"os-schema-check/internal/schema"
	"os-schema-check/internal/validator"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <schema.json> <data file>")
		os.Exit(1)
	}

	// Load and parse OpenSearch Index Schema Settings
	schemaPath := args[0]
	schema, err := schema.Load(schemaPath)
	if err != nil {
		fmt.Printf("Error loading schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parsed JSON: %+v\n", schema)

	// Load and validate bulk data file
	dataPath := args[1]
	reader, err := document.NewDocumentReader(dataPath)
	if err != nil {
		fmt.Printf("Error opening data file: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	invalid := false
	docIndex := 0

	for reader.Next() {
		docIndex++
		doc := reader.Current()
		result := validator.Validate(schema, doc)
		if !result.Valid {
			invalid = true
			fmt.Printf("Document %d is invalid:\n", docIndex)
			for _, issue := range result.Issues {
				fmt.Printf("  - %s: %s\n", issue.Field, issue.Message)
			}
			continue
		}
		fmt.Printf("Document %d is valid\n", docIndex)
	}

	if err := reader.Err(); err != nil {
		fmt.Printf("Error reading data file: %v\n", err)
		os.Exit(1)
	}

	if invalid {
		os.Exit(1)
	}
}
