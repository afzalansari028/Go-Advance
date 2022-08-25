package main

import (
	"fmt"
	"sync"
)

func main() {
	vals := []int{100, 50, 20, 90}
	in := gen(vals)

	// fan-out
	fo1 := square(in)
	fo2 := square(in)

	// fan-in
	fi := merge(fo1, fo2)
	for res := range fi {
		fmt.Println(res)
	}
}

func gen(vals []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, val := range vals {
			out <- val
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func merge(fo ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	fi := func(ch <-chan int) {
		for val := range ch {
			out <- val
		}
		wg.Done()
	}

	wg.Add(len(fo))
	for _, ch := range fo {
		go fi(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
