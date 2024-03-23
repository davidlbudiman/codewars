package katas

import "fmt"

func SeriesSum(n int) string {
	var sum float64
	var i = 1.0
	for j := 0; j < n; j++ {
		sum += 1 / i
		i = i + 3
	}
	return fmt.Sprintf("%.2f", sum)
}
