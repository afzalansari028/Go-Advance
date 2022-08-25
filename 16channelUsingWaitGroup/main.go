package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(3)

	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
		ch <- 1
	}()

	go func() {
		defer wg.Done()
		ch <- 2
	}()

	go func() {
		defer wg.Done()
		ch <- 3
	}()

	// wg.Wait()
	for val := range ch {
		fmt.Printf("value recieved: %v\n", val)
	}
	val, ok := <-ch
	fmt.Println(val, ok)

}
