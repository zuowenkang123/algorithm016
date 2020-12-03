package Week_10

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/decode-string/
// 2020-12-03

func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

// 递归
func decodeString1(s string) string {
	_, str := dnf(s, 0)
	return str
}

func dnf(s string, i int) (int, string) {
	var mul int
	var str string
	var n int
	var temp string
	for n = i; n < len(s); n++ {
		if s[n] >= '0' && s[n] <= '9' {
			mul = mul*10 + int(s[n]-'0')
		} else if s[n] == '[' {
			n, temp = dnf(s, n+1)
			for num := 0; num < mul; num++ {
				str = str + temp
			}
			mul = 0
		} else if s[n] == ']' {
			return n, str
		} else {
			str = str + string(s[n])
		}
	}
	return n, str
}
