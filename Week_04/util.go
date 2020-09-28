package Week_04

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 一步达到
func oneDiff(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

func indexOf(str string, bank []string) int {
	for i, s := range bank {
		if str == s {
			return i
		}
	}
	return -1
}
