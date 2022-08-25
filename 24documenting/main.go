package main

import (
	"fmt"

	"example.local/calc/math"
	"example.local/calc/math/stats"
)

func main() {

	add := math.Add(3, 5)
	fmt.Println("Adition :", add)

	sub := math.Sub(3, 1)
	fmt.Println("Adition :", sub)

	sl := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	avg := stats.Avg(sl)
	fmt.Println("average :", avg)

}
