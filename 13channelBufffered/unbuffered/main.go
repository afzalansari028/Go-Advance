package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go send(ch)
	fmt.Println("\nblocking send goroutine...")
	fmt.Printf("channel length: %v\n", len(ch))

	go recieve(ch)

	time.Sleep(2 * time.Second)
}

func send(ch chan string) {
	ch <- "message"
}

func recieve(ch chan string) {
	time.Sleep(time.Second * 1)
	fmt.Println("send goroutine unblocked")
	fmt.Printf("channel lenght: %v\n", len(ch))
	fmt.Printf("recieved: %v\n", <-ch)
}
