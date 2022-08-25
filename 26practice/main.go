package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(2)
	go send(ch)

	fmt.Println("\nblocking send goroutine...")
	fmt.Printf("channel length: %v\n", len(ch))

	go recieve(ch)
	wg.Wait()
	// time.Sleep(2 * time.Second)
}

func send(ch chan string) {
	ch <- "message"
}

func recieve(ch chan string) {
	// time.Sleep(time.Second * 1)
	var wg sync.WaitGroup
	defer wg.Done()
	fmt.Println("send goroutine unblocked")
	fmt.Printf("channel lenght: %v\n", len(ch))
	fmt.Printf("recieved: %v\n", <-ch)
}
