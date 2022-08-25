package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	st := time.Since(t)
	fmt.Println(st)
}
