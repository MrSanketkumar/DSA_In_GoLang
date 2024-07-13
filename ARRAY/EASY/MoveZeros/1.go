package main

import "fmt"

func main() {
	inputslice := []int{1, 0, 0, 3, 0} //[1,3,12,0,0]
	fmt.Println("Before moving the zero from the slice",inputslice)
	moveZero(inputslice)
	fmt.Println("\nAfter moving the zero from the slice",inputslice)
}

func moveZero(inputSlice []int) {
    lastNonZeroFoundAt := 0
    for i := 0; i < len(inputSlice); i++ {
        if inputSlice[i] != 0 {
            inputSlice[lastNonZeroFoundAt], inputSlice[i] = inputSlice[i], inputSlice[lastNonZeroFoundAt]
            lastNonZeroFoundAt++
        }
    }
}
