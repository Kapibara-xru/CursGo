package main

import (
	"fmt"
)

func Solution(A []int) int {
	B := make([]bool, len(A)+2)
	var g int
	B[0] = true
	for _, c := range A {
		B[c] = true
	}
	for i, c := range B {
		if !c {
			g = i
			break
		}
	}
	return g
}

func main() {
	A := []int{1, 2, 3, 11, 6, 7, 4, 9, 12, 5, 10}
	for _, b := range A {
		fmt.Printf("%d ", b)
	}
	fmt.Printf("\n %d \n", Solution(A))
}
