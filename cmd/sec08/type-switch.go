package main 
 
import "fmt"

func anything(a interface{}) {
    // fmt.Println(a)
    switch v := a.(type) {
    case string:
        fmt.Println(v + " v is type of string")
    case int:
        fmt.Println(v)
        fmt.Println("v is type of integer")
    }

}

func main() {
    anything("aaa")
    anything(123)


    var x interface{} = 3
    i, isInt := x.(int)
    fmt.Println(i + 2, isInt)

    f, isFloat64 := x.(float64)
    fmt.Println(f, isFloat64)

    if x == nil {
        fmt.Println("何もない")
    } else if i, isInt := x.(int); isInt {
        fmt.Println(i, "x is int")
    } else if s, isString := x.(string); isString {
        fmt.Println(s)
    } else {
        fmt.Println("対応していない型")
    }

    switch x.(type) {
    case bool:
        fmt.Println("bool")
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    default:
        fmt.Println("I don't Know")
    }

    switch v := x.(type) {
    case bool:
        fmt.Println(v, "bool")
    case int:
        fmt.Println(v, "int")
    case string:
        fmt.Println(v, "string")
    default:
        fmt.Println("I don't Know")
    }
}