package main

import (
	"fmt"
	"reflect"
)

func main() {
	anon := func(msg string) bool {
		fmt.Println(msg)
		return true
	}
	res := anon("Hello anonymous function !")
	fmt.Println(res)
	fmt.Println("anon is of type: ", reflect.TypeOf(anon))
}
