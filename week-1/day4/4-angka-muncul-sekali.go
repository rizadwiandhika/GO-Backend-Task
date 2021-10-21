package main

import "fmt"

func munculSekali(angka string) []int {
	memo := make(map[rune]int)
	result := make([]int, 0)

	for _, digit := range angka {
		memo[digit]++
	}

	for number, occurrence := range memo {
		if occurrence == 1 {
			value := int(number - 48) // adjustment due to ascii code
			result = append(result, value)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6, 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]
}
