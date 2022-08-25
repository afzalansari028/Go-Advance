package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go myfunc(wg, i)
	}
	// time.Sleep(10 * time.Millisecond)
	wg.Wait()
	fmt.Println("Each goutoutine has run to complete, thanks for waiting!")
}

func myfunc(wg sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("finished executing iteration", i)
}
