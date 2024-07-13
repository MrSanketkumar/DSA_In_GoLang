package main

import "fmt"

func main() {
	inputSlice := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(inputSlice))
}

func singleNumber(nums []int) int {
	seen := make(map[int]int)
	result := 0
	for _, num := range nums {
		seen[num]++
	}

	for i, v := range seen {
		if v == 1 {
			result = i
		}
	}
	return result
}