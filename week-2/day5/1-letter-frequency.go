package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func task1(c chan<- map[rune]int, str string) {
	letterFrequency := make(map[rune]int)
	for _, char := range str {
		letterFrequency[char]++
	}
	fmt.Println("task1 finished")
	c <- letterFrequency
}

func task2(c chan<- map[rune]int, str string) {
	letterFrequency := make(map[rune]int)
	for _, char := range str {
		letterFrequency[char]++
	}
	fmt.Println("task2 finished")
	c <- letterFrequency
}

func main() {
	in := bufio.NewReader(os.Stdin)
	letterFrequency := make(map[rune]int)
	sentence := ""
	channel := make(chan map[rune]int)

	sentence, err := in.ReadString('\n')
	if err != nil {
		log.Fatalln(err.Error())
	}
	sentence = strings.Replace(sentence, "\n", "", 1)

	go task1(channel, sentence[:len(sentence)/2])
	go task2(channel, sentence[len(sentence)/2:])

	time.Sleep(10 * time.Second)

	data1 := <-channel
	data2 := <-channel

	for char, occurence := range data1 {
		letterFrequency[char] += occurence
	}
	for char, occurence := range data2 {
		letterFrequency[char] += occurence
	}

	for char, occurence := range letterFrequency {
		fmt.Printf("%c: %v\n", char, occurence)
	}

}
