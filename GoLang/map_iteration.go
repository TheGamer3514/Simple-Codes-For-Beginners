/*
map_iteration.go
This script demonstrates iterating over a map in Go.
*/

package main

import "fmt"

func main() {
    person := map[string]string{
        "name": "Alice",
        "city": "London",
    }

    for key, value := range person {
        fmt.Printf("%s: %s\n", key, value)
    }
}