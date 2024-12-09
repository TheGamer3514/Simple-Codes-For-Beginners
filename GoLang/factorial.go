/*
factorial.go
This script demonstrates calculating the factorial of a number using recursion.
*/

package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    num := 5
    fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}