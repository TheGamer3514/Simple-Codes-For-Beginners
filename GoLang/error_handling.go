/*
error_handling.go
This script demonstrates basic error handling in Go.
*/

package main

import (
    "fmt"
    "strconv"
)

func main() {
    input := "123a" // Example input string
    if num, err := strconv.Atoi(input); err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted number:", num)
    }
}