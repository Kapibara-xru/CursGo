package main

import (
	"fmt"
)

func Solution(A []int) int {
	flag := 1
	for i := 1; i <= len(A); i++ {
		if !checkElem(A, i) {
			flag = 0
		}
	}
	return flag
}

func checkElem(A []int, n int) bool {
	var flag bool
	for _, c := range A {
		if c == n {
			flag = true
			break
		}
	}
	return flag
}

func main() {
	A := []int{1, 2, 3, 8, 6, 7, 4, 9, 10, 5}
	for i := range A {
		fmt.Printf("%d ", A[i])
	}
	fmt.Printf("\n %d \n", Solution(A))
}
