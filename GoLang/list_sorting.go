/*
list_sorting.go
This script demonstrates sorting a slice in Go.
*/

package main

import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{42, 17, 23, 68, 34}
    fmt.Println("Original slice:", numbers)
    sort.Ints(numbers)
    fmt.Println("Sorted slice (ascending):", numbers)
}