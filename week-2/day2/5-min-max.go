package main

import (
	"fmt"
	"math"
)

func FindMinAndMax(arr []int) string {
	min := math.MaxInt64
	max := math.MinInt64
	indexMin := 0
	indexMax := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			indexMax = i
		}
		if arr[i] < min {
			min = arr[i]
			indexMin = i
		}
	}
	return fmt.Sprintf("min: %d index: %d max: %d index: %d", min, indexMin, max, indexMax)
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}
