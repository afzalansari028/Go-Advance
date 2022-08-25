package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(3 * time.Second)
			ch1 <- i
			time.Sleep(1 * time.Second)
			ch2 <- i
		}
	}()

	for i := 0; i < 4; i++ {
		select {
		case val := <-ch1:
			fmt.Printf("value recieved from channel ch1: %v\n", val)
		case val := <-ch2:
			fmt.Printf("value recieved from channel ch2: %v\n", val)
		default:
			fmt.Println("no data recieved...performing some other operation")
		}
		time.Sleep(2 * time.Second)
	}

}
