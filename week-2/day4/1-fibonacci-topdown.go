package main

import "fmt"

func topdown(memo []int, n int) int {
	if n <= 1 {
		memo[n] = n
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n-1] = topdown(memo, n-1)
	memo[n-2] = topdown(memo, n-2)
	return memo[n-1] + memo[n-2]
}

func fibo(n int) int {
	if n < 0 {
		return -1
	}

	memo := make([]int, n+1)
	return topdown(memo, n)
}

func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55
}
