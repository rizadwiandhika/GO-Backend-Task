package main

import (
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int) {
	check := [1000]int{}

	for x := 2; x < len(check); x++ {
		if x == 2 {
			check[x] = 1
			for y := x + x; y < len(check); y += x {
				check[y] = 2
			}
		} else {
			if check[x] == 0 {
				limit := int(math.Floor(math.Sqrt(float64(x))))
				prima := true
				for z := 2; z <= limit; z++ {
					if x%z == 0 {
						prima = false
						break
					}
				}
				if prima {
					check[x] = 1
					for j := x + x; j < len(check); j += x {
						check[j] = 2
					}
				}
			}
		}
	}
	start++
	baris := 0
	kolom := 0
	total := 0
	for kolom < wide {
		for baris < high {
			if check[start] == 1 {
				total += start
				fmt.Printf("%d ", start)
				baris++
			}
			start++
		}

		fmt.Println()
		baris = 0
		kolom++
	}
	fmt.Println(total)
	fmt.Println()
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}
