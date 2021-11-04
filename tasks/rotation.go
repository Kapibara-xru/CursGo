package tasks

// Task rotation
func Solution1(A []int, k int) []int {
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
