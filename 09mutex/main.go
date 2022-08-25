package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter int
	var wg sync.WaitGroup
	var mut sync.Mutex

	inc := func() {
		mut.Lock()
		defer mut.Unlock()
		counter++
		fmt.Println("increment counter =", counter)
	}

	dec := func() {
		mut.Lock()
		defer mut.Unlock()
		counter--
		fmt.Println("decrement counter =", counter)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dec()
		}()
	}

	wg.Wait()
	fmt.Println("final value of counter", counter)
}
