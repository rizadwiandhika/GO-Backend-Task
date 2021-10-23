package main

import (
	"bufio"
	"fmt"
	"os"
)

type student struct {
	name       string
	nameEncode string

	score int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	encoded := []rune(s.name)
	for i := range encoded {
		encoded[i] = cipherEncode[encoded[i]]
	}

	nameEncode = string(encoded)
	return nameEncode

}

func (s *student) Decode() string {
	var nameDecode = ""

	decoded := []rune(s.nameEncode)
	for i := range decoded {
		decoded[i] = cipherDecode[decoded[i]]
	}

	nameDecode = string(decoded)
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student’s Name : ")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		a.name = line[0 : len(line)-1]

		fmt.Println("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student’s Encode Name : ")

		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		a.nameEncode = line[0 : len(line)-1]

		fmt.Println("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}

var cipherEncode = map[rune]rune{
	' ': ' ',

	'q': 'w',
	'w': 'e',
	'e': 'r',
	'r': 't',
	't': 'y',
	'y': 'u',
	'u': 'i',
	'i': 'o',
	'o': 'p',
	'p': 'a',
	'a': 's',
	's': 'd',
	'd': 'f',
	'f': 'g',
	'g': 'h',
	'h': 'j',
	'j': 'k',
	'k': 'l',
	'l': 'z',
	'z': 'x',
	'x': 'c',
	'c': 'v',
	'v': 'b',
	'b': 'n',
	'n': 'm',
	'm': 'q',
}

var cipherDecode = map[rune]rune{
	' ': ' ',

	'q': 'm',
	'w': 'q',
	'e': 'w',
	'r': 'e',
	't': 'r',
	'y': 't',
	'u': 'y',
	'i': 'u',
	'o': 'i',
	'p': 'o',
	'a': 'p',
	's': 'a',
	'd': 's',
	'f': 'd',
	'g': 'f',
	'h': 'g',
	'j': 'h',
	'k': 'j',
	'l': 'k',
	'z': 'l',
	'x': 'z',
	'c': 'x',
	'v': 'c',
	'b': 'v',
	'n': 'b',
	'm': 'n',
}
