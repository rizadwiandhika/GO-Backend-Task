package main

import "fmt"

type Info struct {
	index, occurrence int
}

func PairSum(arr []int, target int) []int {
	mapping := make(map[int]*Info)

	for index, value := range arr {
		_, isValueExist := mapping[value]

		if isValueExist == false {
			mapping[value] = &Info{index: index, occurrence: 1}
			continue
		}

		mapping[value].occurrence += 1
	}

	for _, value := range arr {
		pairValue := target - value
		_, isPairValueExist := mapping[pairValue]

		if isPairValueExist && pairValue == value && mapping[pairValue].occurrence == 1 {
			continue
		}

		if isPairValueExist && pairValue == value {
			return []int{mapping[pairValue].index, mapping[pairValue].index + 1}
		}
		if isPairValueExist {
			return []int{mapping[value].index, mapping[pairValue].index}
		}
	}

	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 7, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]
}
