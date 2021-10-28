package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return n
	}

	memo := make([]int, n+1)

	// base case for fibonacci
	memo[0] = 0
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
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
