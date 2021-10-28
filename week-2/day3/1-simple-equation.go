package main

import "fmt"

func eq1(x, y, z, a int) bool {
	return x+y+z == a
}
func eq2(x, y, z, b int) bool {
	return x*y*z == b
}
func eq3(x, y, z, c int) bool {
	return x*x+y*y+z*z == c
}
func isSatisfyAllEquation(x, y, z, a, b, c int) bool {
	return eq1(x, y, z, a) && eq2(x, y, z, b) && eq3(x, y, z, c)
}

func simpleEquation(a int, b int, c int) {
	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000; y++ {
			for z := 1; z <= 10000; z++ {
				if isSatisfyAllEquation(x, y, z, a, b, c) {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}

	fmt.Println("no solution")
}

func main() {
	simpleEquation(6, 6, 14)
}
