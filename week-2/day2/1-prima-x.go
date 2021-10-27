package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	check := [1000]int{}

	for i := 2; i < len(check); i++ {
		if i == 2 {
			check[i] = 1
			for j := i + i; j < len(check); j += i {
				check[j] = 2
			}
		} else {
			if check[i] == 0 {
				batas := int(math.Floor(math.Sqrt(float64(i))))
				prime := true
				for k := 2; k <= batas; k++ {
					if i%k == 0 {
						prime = false
						break
					}
				}
				if prime {
					check[i] = 1
					for j := i + i; j < len(check); j += i {
						check[j] = 2
					}
				}
			}
		}
	}

	index := 2
	counter := 0
	result := 0

	for counter < number {
		if check[index] == 1 {
			counter++
			result = index
		}
		index++
	}

	return result
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
