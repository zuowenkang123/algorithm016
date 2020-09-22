package Week_03

import "strings"

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
// 2020-09-22
func replaceSpace(s string) string {
	var res strings.Builder
	for i := range s {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}
