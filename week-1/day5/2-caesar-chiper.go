package main

import "fmt"

func caesar(offset int, str string) string {
	shiftedString := []rune(str)

	for index, char := range str {
		// Dont shift space character
		if char == ' ' {
			continue
		}

		shift := char + rune(offset)%26
		if shift > 'z' {
			start := 'a' - 1
			shift = start + (shift - 'z')
		}
		shiftedString[index] = shift
	}

	return string(shiftedString)
}

func main() {
	fmt.Println(caesar(26, "abc zbds"))
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
