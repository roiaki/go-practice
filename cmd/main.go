package main

import (
    "fmt"
    // "time"
)

func outer() {
    var s4 string ="outer"
    fmt.Println(s4)
}

func sec03() {
    fmt.Println("Hello golang from docker!")
    fmt.Println(time.Now())
    var i int = 100
    fmt.Println(i)

    var s string = "hello Go"
    fmt.Println(s)

    var t, f bool = true, false
    fmt.Println(t, f)

    var (
        i2 int = 200
        s2 string = "Golang" 
    )
    fmt.Println(i2, s2)
    
    var i3 int
    var s3 string
    fmt.Println(i3, s3)

    i3 = 300
    s3 = "Go"
    fmt.Println(i3, s3)

    i = 150
    fmt.Println(i)

    // 暗黙的な定義
    i4 := 400
    fmt.Println(i4)
}

func sec04() {
    var s5 string = "not use"
    var s6 string = s5
    fmt.Println(s6)

    var i1 int = 100
    var i2 int64 = 200
    fmt.Println(i1, i2)
    outer()

    va
}

func main() {
    

    sec04()

    

}

