package Week_10

import (
	"math"
	"strings"
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/
// 2020-12-02

func myAtoi(s string) int {
	//去掉收尾空格
	s = strings.TrimSpace(s)
	result := 0
	sign := 1

	for i, v := range s {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		// 数值最大检测
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * result
}
