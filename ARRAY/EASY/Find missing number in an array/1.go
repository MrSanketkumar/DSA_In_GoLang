package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	// Output: 2
	fmt.Println("The missing number from given sequance is :",missingNumber(nums))
}

func missingNumber(inputSlice []int) int {
	number := len(inputSlice) * (len(inputSlice) + 1) / 2  // find the natural number 
	sum := 0
	for _, v := range inputSlice {
		sum += v
	}
	return number - sum  //natural number -sum of all the number => gives missing number from the slice
}
