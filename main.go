package main

import (
    "fmt"
    "os"
    "strings"
)

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

    if _, err := os.Stat(path); err == nil {
        fmt.Println("File exists")
    } else if os.IsNotExist(err) {
        fmt.Println("File does not exist")
        os.Exit(1)
    } else {
        fmt.Printf("Error checking path: %v\n", err)
        os.Exit(1)
    }
}
