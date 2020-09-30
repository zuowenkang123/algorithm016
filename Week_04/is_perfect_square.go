package Week_04

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 1, num/2

	for l < r {
		mid := (l+r)/2 + 1
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			l = mid

		} else {
			r = mid - 1
		}
	}
	return false
}
