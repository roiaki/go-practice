package main

import (
    "fmt"
    "time"
)

func reciever(c chan int) {
    for {
        i := <-c
        fmt.Println(i)
    }
}

func main() {
    var ch1 chan int
   
    ch1 = make(chan int)

    ch1 <- 1

    go reciever(ch1)
    go reciever(ch2)

    i := 0
    for i < 100 {
        ch1 <- i
        ch2 <- i
        //
        time.Sleep(50 * time.Millisecond)
        i++
    }
}
