package stats

//this function can find the average of slice values which will be passed
func Avg(vals []float64) float64 {
	total := 0.0
	for _, val := range vals {
		total += val
	}
	return total / float64(len(vals))
}
