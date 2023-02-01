package main

import "fmt"

func main() {
    sl := []int{100, 200}
    sl2 := sl

    sl2[0] = 1000
    fmt.Println(sl)
    fmt.Println(sl2)

}