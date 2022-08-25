package main

import (
	"fmt"
	"reflect"
)

func main() {

	myfunc := retfun()
	fmt.Println(myfunc())
	fmt.Println("myfunc is of type: ", reflect.TypeOf(myfunc))
}

func retfun() func() bool {
	return func() bool {
		return true
	}
}
