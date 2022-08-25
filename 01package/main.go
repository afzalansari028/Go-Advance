package main

import (
	"fmt"
	//    m "example.local/package/math"
	"example.local/package/math"
	"example.local/package/math/stats"
)

func main() {

	vals := []float64{1, 10, 100, 1000}
	fmt.Println(math.Add(20, 10))
	fmt.Println(stats.Avg(vals))

}
