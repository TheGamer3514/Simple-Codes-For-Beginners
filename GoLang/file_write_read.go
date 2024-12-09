/*
file_write_read.go
This script demonstrates writing to and reading from a file in Go.
*/

package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    filename := "example.txt"
    content := []byte("Hello, File!")
    
    // Write to file
    if err := ioutil.WriteFile(filename, content, 0644); err != nil {
        log.Fatal(err)
    }
    fmt.Println("File written successfully")

    // Read from file
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File content:", string(data))
}