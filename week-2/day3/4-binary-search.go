package main

import "fmt"

func BinarySearch(array []int, x int) int {
	low := 0
	high := len(array) - 1

	if high == 0 {
		return array[high]
	}

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == x {
			return mid
		}

		if array[mid] < x {
			low = mid + 1
		}
		if x < array[mid] {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))                  // 2)
	fmt.Println(BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5))                 // 3)
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))  // 6)
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)) // -1)
}
