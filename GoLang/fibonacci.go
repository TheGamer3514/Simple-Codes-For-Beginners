/*
fibonacci.go
This script demonstrates generating the Fibonacci sequence.
*/

package main

import "fmt"

func main() {
    terms := 10
    a, b := 0, 1
    fmt.Println("Fibonacci sequence:")
    for i := 0; i < terms; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()
}