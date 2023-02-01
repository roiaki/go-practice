package main

import "fmt"

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for index, value := range sl {
		fmt.Println(index, value)
	}

	for _, value := range sl {
		fmt.Println(value)
	}
}