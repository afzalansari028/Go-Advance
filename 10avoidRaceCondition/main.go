package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter int64
	var wg sync.WaitGroup
	// var mut sync.Mutex

	inc := func() {
		for i := 0; i < 1000; i++ {
			// mut.Lock()
			counter++
			// mut.Unlock()
		}
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
