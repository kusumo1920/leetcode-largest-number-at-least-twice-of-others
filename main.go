package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{0, 0, 3, 2}
	output := dominantIndex(input)
	fmt.Println(output)
}

func dominantIndex(nums []int) int {
	doubleLargest := math.MinInt
	secondDoubleLargest := math.MinInt
	singleLargest := math.MinInt
	singleLargestIndex := -1

	for i, v := range nums {
		if v >= singleLargest {
			singleLargest = v
			singleLargestIndex = i
		}
		if v*2 > doubleLargest {
			secondDoubleLargest = doubleLargest
			doubleLargest = v * 2
		} else if v*2 > secondDoubleLargest {
			secondDoubleLargest = v * 2
		}
	}

	if singleLargest < secondDoubleLargest {
		return -1
	}

	return singleLargestIndex
}
