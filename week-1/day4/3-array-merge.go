package main

import (
	"fmt"
	"math"
)

func ArrayMerge(arrA, arrB []string) []string {
	lenA := len(arrA)
	lenB := len(arrB)
	maxLength := int(math.Max(float64(lenA), float64(lenB)))
	minLength := int(math.Min(float64(lenA), float64(lenB)))

	memo := make(map[string]bool)
	merged := make([]string, 0, int(maxLength)) // min length of merged arrays

	for i := 0; i < minLength; i++ {
		value1 := arrA[i]
		value2 := arrB[i]

		if memo[value1] && memo[value2] {
			continue
		}

		if memo[value1] {
			memo[value2] = true
			merged = append(merged, value2)
			continue
		}

		if memo[value2] {
			memo[value1] = true
			merged = append(merged, value1)
			continue
		}

		if value1 == value2 {
			memo[value1] = true
			merged = append(merged, value1)
			continue
		}

		memo[value1] = true
		memo[value2] = true
		merged = append(merged, value1, value2)
	}

	for ; minLength < maxLength; minLength++ {
		var value string

		if lenA > minLength {
			value = arrA[minLength]
		}
		if lenB > minLength {
			value = arrB[minLength]
		}

		if memo[value] {
			continue
		}

		memo[value] = true
		merged = append(merged, value)
	}

	return merged
}

func main() {
	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// [“king”, “devil jin”, “akuma”, “eddie”, “steve”, “geese”]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
