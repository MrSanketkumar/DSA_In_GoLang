package main

import "fmt"

func main() {
	inputSlice := []int{1,1,0,1,1,1}
	fmt.Println("The maximum Consecutive onces is :",findMaxConsecutiveOnes(inputSlice)) //3
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}