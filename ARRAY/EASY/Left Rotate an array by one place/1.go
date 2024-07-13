package main

import "fmt"

func main() {
	inputslice := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Println("Before the left rotate :", inputslice)
	leftRotate(inputslice, k) //[5,6,7,1,2,3,4]
	fmt.Println("\nAfter the left rotate :", inputslice)
}

func leftRotate(inputslice []int, k int) {
	k = k % len(inputslice) // to make the k fit in the range of len of the array or slice if k>len(inputslice)
	reverse(inputslice[0 : len(inputslice)-k])
	reverse(inputslice[k+1:])
	reverse(inputslice[0:])
}

func reverse(inputslice []int) {
	i := 0
	j := len(inputslice) - 1
	for i < j {
		inputslice[i], inputslice[j] = inputslice[j], inputslice[i]
		i++
		j--
	}
}
