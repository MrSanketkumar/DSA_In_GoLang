package main

import "fmt"

func main() {
	inputSlice := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicate(inputSlice))
}

func removeDuplicate(inputSlice []int) []int {
	result := []int{}
	seen := make(map[int]bool)
	for _, value := range inputSlice {
		if seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}
