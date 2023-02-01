package main

import "fmt"

func main() {
    var m = map[string]int{"A": 100, "B": 200}

    fmt.Println(m)

    m2 := map[string]int{"A": 100, "B": 200}
    fmt.Println(m2)

    m3 := map[string]int{
        "A": 100, 
        "B": 200,
    }
    fmt.Println(m3)

    m4 := make(map[int]string)
    fmt.Println(m4)

    m4[1] = "Japan"
    m4[2] = "Usa"
    fmt.Println(m4)

    fmt.Println(m["A"])

    s, doHave := m4[4]
    fmt.Println(s, doHave)
    if !doHave {
        fmt.Println("not have")
    }

    m4[3] = "ABC"
    fmt.Println(m4)

    delete(m4, 3)
    fmt.Println(m4)

    m5 := map[string]int {
        "a": 100,
        "b": 200,
    }

    for key, value := range m5 {
        fmt.Println(key, value)
    }
}