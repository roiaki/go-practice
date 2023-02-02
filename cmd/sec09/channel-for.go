package main

import (
    "fmt"
)

func main() {
    ch1 := make(chan int, 2)
    ch2 := make(chan string, 2)
    ch2 <- "A"
   /*
 

   v1 := <-ch1
   v2 := <-ch2
   fmt.Println(v1)
   fmt.Println(v2)
   */
    select {
    case v1 := <-ch1:
        fmt.Println(v1 + 100)
    case v2 := <-ch2:
        fmt.Println(v2 + " v2")
    }

    ch3 := make(chan int)
    ch4 := make(chan int)
    ch5 := make(chan int)

    // reciever
    go func() {
        for {
            i := <- ch3
            ch4 <- i * 2
        }
    }()

    go func() {
        for {
            i2 := <- ch4
            ch5 <- i2 - 1
        }
    }()
}