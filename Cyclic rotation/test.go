package main

import (
	"fmt"
)

func Solution(A []int, k int) []int {
	var j, tmp, n int
	n = len(A) - 1
	for j < k {
		tmp = A[n]
		for i := n; i > 0; i-- {
			A[i] = A[i-1]
		}
		A[0] = tmp
		j++
	}
	return A
}

func main() {
	A := []int{8, 5, 7, 6}
	for _, b := range A {
		fmt.Printf("%d ", b)
	}
	A = Solution(A, 3)
	fmt.Printf("\n")
	for _, b := range A {
		fmt.Printf("%d ", b)
	}
}
