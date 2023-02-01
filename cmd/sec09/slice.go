package main

import "fmt"

func main() {
	var sl1 []int
	fmt.Println(sl1)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	var sl3 = []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:3])
	fmt.Println(sl5[4:])

	fmt.Println(len(sl5))

}