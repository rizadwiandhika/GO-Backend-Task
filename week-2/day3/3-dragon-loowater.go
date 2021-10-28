package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead []int, knightHeight []int) {
	totalDragon := len(dragonHead)
	totalKnight := len(knightHeight)

	// Dragon always win if total dragon are more than knight
	if totalDragon > totalKnight {
		fmt.Println("knight failed")
		return
	}

	// Two ways of sorting: (complexity: nlog(n))
	sort.Ints(knightHeight)
	sort.Slice(dragonHead, func(i, j int) bool { return dragonHead[i] < dragonHead[j] })

	totalHeight := 0
	d, k := 0, 0

	for d < totalDragon && k < totalKnight {
		dragon := dragonHead[d]
		knight := knightHeight[k]

		if dragon <= knight {
			totalHeight += knight
			d++
		}

		k++
	}

	// All dragons are defeated
	if d == totalDragon {
		fmt.Println(totalHeight)
		return
	}

	fmt.Println("knight failed")
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
