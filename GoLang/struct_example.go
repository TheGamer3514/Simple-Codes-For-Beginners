/*
struct_example.go
This script demonstrates defining and using a struct in Go.
*/

package main

import "fmt"

type Dog struct {
    Name  string
    Breed string
}

func (d Dog) Bark() {
    fmt.Printf("%s says Woof!\n", d.Name)
}

func main() {
    myDog := Dog{"Buddy", "Golden Retriever"}
    myDog.Bark()
}