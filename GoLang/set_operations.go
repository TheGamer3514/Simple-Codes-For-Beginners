/*
set_operations.go
This script demonstrates set-like operations using maps in Go.
*/

package main

import "fmt"

func main() {
    set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
    set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

    // Intersection
    intersection := map[int]bool{}
    for key := range set1 {
        if set2[key] {
            intersection[key] = true
        }
    }
    fmt.Println("Intersection:", intersection)
}