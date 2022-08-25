package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := inc()
	fmt.Println("x is of type: ", reflect.TypeOf(x))

	for i := 0; i < 4; i++ {
		fmt.Println("x = ", x())
	}

	y := inc()
	fmt.Println("y = ", y())
}

func inc() func() int {
	var j int
	return func() int {
		j++
		return j
	}

}
