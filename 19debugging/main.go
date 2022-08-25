package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Printf("addding %v and %v\n", x, y)
	fmt.Println("result from operation: ", Add(x, y))
}
func Add(x, y int) int {
	res := x + y
	// res = res + 1
	return res
}
