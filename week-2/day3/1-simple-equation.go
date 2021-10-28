package main

var A int
var B int
var C int

func df(x float64) float64 {
	a := float64(A)
	return 3*x*x - 2*a*x + 0.5*a*a
}

func f(x float64) float64 {
	a := float64(A)
	b := float64(B)
	return x*x*x - a*x*x + 0.5*a*a*x - b
}

func findRootNewton() float64 {
	iteration := 100
	root := 100.0

	for i := 0; i < iteration; i++ {
		root = root - f(root)/df(root)
	}

	return root
}

func simpleEquation(a int, b int, c int) {
	A = a
	B = b
	C = c

	x := findRootNewton()
	y := ...
	c := ...
}

func main() {
	simpleEquation(6, 6, 14)
}
