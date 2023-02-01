package main

import (
    "fmt"
)

// const
const Pi = 3.14
const (
	URL 	 = "http://xxx.co.jp"
	SiteName = "test"
)

func main() {
	// 型
	// 浮動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	
	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	// 論理値型
	var t, f bool = true, false
	fmt.Println(t, f)

	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	
	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))

	// byte(uint8)型
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	// 配列型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	var arr5 [5]int = [5]int{1, 2, 3}
	fmt.Println(arr5)

	fmt.Println(arr4[0])
	fmt.Println(arr4[2-1])

	arr4[1] = "F"
	fmt.Println(arr4)

	// interface nil
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	// 型変換
	// fl64 := 10.2

	// n := int(fl64)
	// nn := float64(n)

	// var s string = "100"

	// i, _ := strconv.Atoi(s)
	// fmt.Printf("%T %v\n", i, i)

	// h := "Hello World"
	// fmt.Println(string(h[0]))

	// var ii int = 321
	// var ss string
	// ss = strconv.Itoa(ii)
	// fmt.Println(ss)

	// var s string = "100"
	// i, _ := strconv.Atoi(s)

	// const
	fmt.Println(URL, SiteName)

}