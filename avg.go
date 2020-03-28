package main

import (
	"fmt"
	"math"
)

// variadic function
func avg(s ...int) float64 {
	total := 0
	numbers := len(s)
	for _, num := range s {
		total += num
	}
	avg := float64(total) / float64(numbers)
	return math.Round(avg*100) / 100
}

func main() {
	fmt.Println(avg(1, 2, 3, 4, 7))
}
