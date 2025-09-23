package main

import (
    "fmt"
    "os"
    "os-schema-check/internal/fileutil"
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

    dataPath := args[1]
    if !fileutil.IsAvailable(dataPath) {
        fmt.Printf("Error checking second path: %v\n", dataPath)
        os.Exit(1)
    }

    fmt.Println("Second file exists")
}
