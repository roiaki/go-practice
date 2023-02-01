package main

import (
    "fmt"
    "os"
)

func Testdefer() {
    defer fmt.Println("END")
    fmt.Println("START")
}

func main() {

    file, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    file.Write([]byte("Hello Golang\n"))

}