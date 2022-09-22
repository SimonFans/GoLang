package Basic

func IsOdd(n int) bool {
	return n%2 == 1
}

func Filter(pred func(int) bool, values []int) []int {
	var out []int
	for _, val := range values {
		if pred(val) {
			out = append(out, val)
		}
	}
	return out
}
