package tasks

// Task search for absent
func Solution2(A []int) int {
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
