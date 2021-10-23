package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	n := float64(len(s.score))
	total := 0.0

	for _, score := range s.score {
		total += float64(score)
	}

	return total / n
}

func (s Student) Min() (min int, name string) {
	min = math.MaxInt64

	for index, score := range s.score {
		if score < min {
			min = score
			name = s.name[index]
		}
	}
	return
}

func (s Student) Max() (max int, name string) {
	max = math.MinInt64

	for index, score := range s.score {
		if score > max {
			max = score
			name = s.name[index]
		}
	}
	return
}

func main() {
	var a = Student{}

	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Print("Input " + string(i) + " Student’s Name : ")
		fmt.Scan(&name)

		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)

		a.name = append(a.name, name)
		a.score = append(a.score, score)

	}

	fmt.Printf("%+v", a)

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
