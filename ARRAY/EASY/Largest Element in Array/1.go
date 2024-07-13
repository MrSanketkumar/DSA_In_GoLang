package main

import (
	"fmt"
	"math"
)

func main() {
	inputSlice := []int{10}
	fmt.Println("the max element in the slice is :",largesetNumber(inputSlice))
}

func largesetNumber(inputSlice []int) int {
	maxElement := math.MinInt

	for _, value := range inputSlice {
		if maxElement<value{
			maxElement = value
		}
	}
	return maxElement
}