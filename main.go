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

    fmt.Println(strings.Join(args, " "))
}
