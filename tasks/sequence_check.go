package tasks

// Task sequence check
func Solution3(A []int) int {
	flag := 1
	for i := 1; i <= len(A); i++ {
		if !checkElem(A, i) {
			flag = 0
			break
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
