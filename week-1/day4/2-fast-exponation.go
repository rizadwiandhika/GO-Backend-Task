package main

import "fmt"

func pow(x int, n int) int {
	memo := make([]int, n+1)
	return topDown(x, n, memo)
}

func topDown(x int, n int, memo []int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if memo[n] != 0 {
		return memo[n]
	}

	right := n / 2
	left := n - right

	memo[left] = topDown(x, left, memo)
	memo[right] = topDown(x, right, memo)
	return memo[left] * memo[right]
}

func main() {
	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(7, 2)) // 49

	fmt.Println(pow(10, 5)) // 100000

	fmt.Println(pow(17, 6)) // 24137569

	fmt.Println(pow(5, 3)) // 125
}
