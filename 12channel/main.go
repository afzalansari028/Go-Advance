package main

import "fmt"

func main() {
	ch := make(chan int)
	mult := func(x, y int) {
		res := x * y
		ch <- res
	}
	go mult(10, 10)
	val, ok := <-ch
	fmt.Printf("type of ch is: %T\n", ch)
	fmt.Printf("the value of ch is: %v\n", ch)
	fmt.Printf("the resulting value from the goroutine is: %v\n", val)
	fmt.Printf("the value of ok is: %v", ok)
}
