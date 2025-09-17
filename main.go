package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
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
    if len(args) == 0 {
        fmt.Println("Hello, World!")
        return
    }

    path := args[0]
    if !strings.HasSuffix(path, ".json") {
        fmt.Println("Error: file must have .json extension")
        os.Exit(1)
    }

    if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            fmt.Println("File does not exist")
            os.Exit(1)
        }

        fmt.Printf("Error checking path: %v\n", err)
        os.Exit(1)
    }

    content, err := os.ReadFile(path)
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
}
