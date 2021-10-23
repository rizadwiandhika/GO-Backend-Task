package main

import (
	"fmt"
	"math"
)

func getMinMax(numbers ...*int) (min int, max int) {
	min = math.MaxInt64
	max = math.MinInt64

	for _, vp := range numbers {
		value := *vp

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
