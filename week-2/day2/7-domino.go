package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) string {
	result := []int{-1, -1}
	for i := 0; i < len(cards); i++ {
		if cards[i][0] == deck[0] || cards[i][0] == deck[1] {
			if result[0]+result[1] < cards[i][0]+cards[i][1] {
				result[0] = cards[i][0]
				result[1] = cards[i][1]
			}
		}
		if cards[i][1] == deck[0] || cards[i][1] == deck[1] {
			if result[0]+result[1] < cards[i][0]+cards[i][1] {
				result[0] = cards[i][0]
				result[1] = cards[i][1]
			}
		}
	}

	if result[0] == -1 && result[1] == -1 {
		return "tutup kartu"
	} else {
		return fmt.Sprintf("[%d, %d]", result[0], result[1])
	}
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
