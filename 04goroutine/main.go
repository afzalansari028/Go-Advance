package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("running in main routine")
	go g1()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("executed from main routine")

}

func g1() {
	fmt.Println("running in goroutine g1")
}
