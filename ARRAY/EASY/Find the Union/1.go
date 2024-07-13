package main

import "fmt"

func main() {
	inputSlice1 := []int{1, 2, 3, 4, 5}
	inputSlice2 := []int{1, 2, 3, 6, 7}

	fmt.Println("The union of two slice is :",union(inputSlice1, inputSlice2)) // 1 2 3 4 5 6 7
}

func union(inputSlice1, inputSlice2 []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(inputSlice1) && j < len(inputSlice2) {
		if inputSlice1[i] > inputSlice2[j] {
			result = append(result, inputSlice1[j])
			j++
		} else if inputSlice1[i] < inputSlice2[j] {
			result = append(result, inputSlice1[i])
			i++
		} else {
			result = append(result, inputSlice1[i])
			i++
			j++
		}
	}

	if i<len(inputSlice1){
		result = append(result, inputSlice1[i])
		i++
	}

	if j<len(inputSlice2){
		result = append(result, inputSlice2[j])
		j++
	}
	return result
}
