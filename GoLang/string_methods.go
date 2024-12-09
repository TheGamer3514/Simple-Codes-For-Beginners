/*
string_methods.go
This script demonstrates string manipulation in Go.
*/

package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "Hello, Go!"
    fmt.Println("Original:", text)
    fmt.Println("To Upper:", strings.ToUpper(text))
    fmt.Println("To Lower:", strings.ToLower(text))
    fmt.Println("Replace:", strings.ReplaceAll(text, "Go", "Golang"))
}