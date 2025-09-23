package main

import (
	"fmt"
	"os"
	"os-schema-check/internal/document"
	"os-schema-check/internal/schema"
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

	for reader.Next() {
		fmt.Printf("Document: %+v\n", reader.Current())
	}

	if err := reader.Err(); err != nil {
		fmt.Printf("Error reading data file: %v\n", err)
		os.Exit(1)
	}
}
