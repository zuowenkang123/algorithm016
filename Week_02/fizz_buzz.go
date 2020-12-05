package Week_02

import "strconv"

// https://leetcode-cn.com/problems/fizz-buzz/
// 2020-09-17

func fizzBuzz(n int) []string {
	res := make([]string, 0)
	for i := 1; i <= n; i++ {
		str := ""
		if i%3 != 0 && i%5 != 0 {
			str = strconv.Itoa(i)
		}
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}

		res = append(res, str)
	}
	return res
}
