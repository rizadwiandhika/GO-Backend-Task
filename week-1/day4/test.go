package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := make([]int, len(arr1))

	copy(arr2, arr1)
	arr2[0] = 999

	fmt.Println(arr1, arr2)
}
