package main

import (
    "fmt"
    "time"
)


func main() {

    ch1 := make(chan int)
    ch2 := make(chan int)

    go reciever1(ch1)
    go reciever1(ch2)

    i := 0
    for i < 100 {
        ch1 <- i
        ch2 <- i
        //
        time.Sleep(50 * time.Millisecond)
        i++
    }
}