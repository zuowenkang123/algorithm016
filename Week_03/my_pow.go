package Week_03

// https://leetcode-cn.com/problems/powx-n/
// 2020-09-24

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	r := myPow(x, n/2)
	if n%2 == 0 {
		r *= r
	} else {
		r = r * r * x
	}
	return r
}

func myPow3(x float64, n int) float64 {
	// n大于等于0
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	// 分治
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func myPow1(x float64, n int) float64 {
	if n < 0 {
		x = 1.0 / x
		n = -n
	} else if n == 0 {
		return 1
	}
	ans := 1.0
	for n > 1 {
		if n&1 == 1 {
			ans *= x
			n--
		} else {
			x = x * x
			n /= 2
		}
	}
	ans *= x
	return ans
}

func myPow2(x float64, n int) float64 {
	if n < 0 {
		return help(1.0/x, -n)
	} else if n == 0 {
		return 1
	} else {
		return help(x, n)
	}
}
func help(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n&1 == 0 {
		return help(x*x, n/2)
	} else {
		return x * help(x, n-1)
	}
}
