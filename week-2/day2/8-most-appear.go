package main

import "fmt"

func MostAppearItem(items []string) map[string]int {
	barang := map[string]int{}

	for i := 0; i < len(items); i++ {
		if _, ok := barang[items[i]]; ok {
			barang[items[i]]++
		} else {
			barang[items[i]] = 1
		}
	}
	return barang
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
