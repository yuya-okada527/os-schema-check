package main

import (
    "encoding/json"
    "fmt"
    "os"
    "os-schema-check/internal/fileutil"
)

type Property struct {
    Type string `json:"type"`
}

type Mappings struct {
    Properties map[string]Property `json:"properties"`
}

type Schema struct {
    Mappings Mappings `json:"mappings"`
}

func main() {
    args := os.Args[1:]
    if len(args) != 2 {
        fmt.Println("Usage: go run main.go <schema.json> <data file>")
        os.Exit(1)
    }

    schemaPath := args[0]
    if !fileutil.CheckExtension(schemaPath, ".json") {
        fmt.Println("Error: file must have .json extension")
        os.Exit(1)
    }

    if !fileutil.IsAvailable(schemaPath) {
        fmt.Printf("Error checking path: %v\n", schemaPath)
        os.Exit(1)
    }

    content, err := os.ReadFile(schemaPath)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    var schema Schema
    if err := json.Unmarshal(content, &schema); err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
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
