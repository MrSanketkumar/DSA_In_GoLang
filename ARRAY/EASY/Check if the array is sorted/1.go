package main

import "fmt"

func main() {
	inputSlice := []int{3, 4, 5, 2}
	fmt.Println(check(inputSlice))
}

func check(inputSlice []int) bool {
	var count int
	for i, v := range inputSlice {
		if v > inputSlice[(i+1)%len(inputSlice)] {
			count += 1
		}
		if count > 1 {
			return false
		}
	}
	return true
}