package main

import (
	"fmt"
)

func Solution(A []int) int {
	var elem, count int
	for _, c := range A {
		for _, g := range A {
			if c == g {
				count++
			}
		}
		if count%2 != 0 {
			elem = c
			break
		}
	}
	return elem
}

func main() {
	A := []int{1, 1, 2, 3, 4, 2, 5, 3, 5}
	for i := range A {
		fmt.Printf("%d ", A[i])
	}
	fmt.Printf("\n %d \n", Solution(A))
}
