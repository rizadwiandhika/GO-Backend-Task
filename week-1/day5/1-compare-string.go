package main

import (
	"fmt"
	"strings"
)

func Compare(str1 string, str2 string) string {
	lenStr1 := len(str1)
	lenStr2 := len(str2)

	if lenStr1 > lenStr2 && strings.Contains(str1, str2) {
		return str2
	}
	if strings.Contains(str2, str1) {
		return str1
	}
	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI")) // AKA

	fmt.Println(Compare("KANGOORO", "KANG")) // KANG

	fmt.Println(Compare("KI", "KIJANG")) // KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	fmt.Println(Compare("ILALANG", "ILA")) // ILA
}
