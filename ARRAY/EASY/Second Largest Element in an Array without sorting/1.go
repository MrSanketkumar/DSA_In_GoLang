package main

import (
	"fmt"
	"math"
)

func main() {
	inputSlice := []int{1, 8, 7, 56, 90}
	fmt.Println("The second max element is :",largesetNumber(inputSlice))
}

func largesetNumber(inputSlice []int) int {
	maxElement := math.MinInt
	secondMax := 0
	for _, value := range inputSlice {
		if value > maxElement {
			secondMax = maxElement
			maxElement = value
		}
		if maxElement != value && secondMax < value {
			secondMax = value
		}
	}
	return secondMax
}
