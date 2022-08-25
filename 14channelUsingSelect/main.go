package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(10 * time.Second)
		ch2 <- 2
	}()

	go func() {
		time.Sleep(10 * time.Second)
		ch3 <- 3
	}()

	for i := 0; i < 3; i++ {
		select {
		case val1 := <-ch1:
			fmt.Printf("value recieved from ch1: %v\n", val1)
		case val2 := <-ch2:
			fmt.Printf("value recieved from ch2: %v\n", val2)
		case val3 := <-ch3:
			fmt.Printf("value recieved from ch3: %v\n", val3)
		case to := <-time.After(5 * time.Second):
			fmt.Printf("timed out after 5 seconds at %v\n", to)
		}
	}
}
