package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("Hello, World!")
        return
    }

    path := args[0]
    if _, err := os.Stat(path); err == nil {
        fmt.Println("Yes")
    } else if os.IsNotExist(err) {
        fmt.Println("No")
    } else {
        fmt.Printf("Error checking path: %v\n", err)
    }
}
