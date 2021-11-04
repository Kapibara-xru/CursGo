package tasks

// Task wonderful array entries
func Solution4(A []int) int {
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
