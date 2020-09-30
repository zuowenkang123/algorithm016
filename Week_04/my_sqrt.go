package Week_04

import "math"

// https://leetcode-cn.com/problems/sqrtx/
// 2020-09-30

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	res := int(math.Exp(0.5 * math.Log(float64(x))))
	if (res+1)*(res+1) <= x {
		return res + 1
	}
	return res
}

// 二分查找法
func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x/2
	for l < r {
		mid := (r+l)/2 + 1
		if mid*mid > x {
			r = mid - 1

		} else {
			l = mid
		}
	}
	return l
}

func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
