/*
lambda_like.go
This script demonstrates the use of anonymous functions in Go.
*/

package main

import "fmt"

func main() {
    square := func(x int) int {
        return x * x
    }
    fmt.Println("Square of 5:", square(5))
}