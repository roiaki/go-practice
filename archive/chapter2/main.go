package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {    
    flag.Parse()
    arg := flag.Arg(0)
    msg := fmt.Sprintf("Hello %s\n", arg)

    file, err := os.Create("hello.txt")
    if err != nil {
        fmt.Printf("failed to create file \n: %v", err)
        return
    }
    defer file.Close()

    _, err = file.WriteString(msg)
    if err != nil {
        fmt.Printf("failed to write message to file \n: %v", err)
        return
    }

    fmt.Println("Done!")
}