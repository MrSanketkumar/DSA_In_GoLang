package main

import "fmt"

func main() {
	inputSlice := []int{1, 2, 3, 4, 6}
	target := 60
	fmt.Println(linearSearch(inputSlice, target))

}

func linearSearch(inputSlice []int, target int) int {

	for i := 0; i < len(inputSlice); i++ {
		if target == inputSlice[i] {
			return i
		}
	}

	return -1
}
