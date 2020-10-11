package Week_04

import "sort"

// https://leetcode-cn.com/problems/assign-cookies/description/
// 2020-10-11

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	res := 0
	for i := 0; i < len(g); i++ {
		for j < len(s) {
			if g[i] <= s[j] {
				res++
				j++
				break
			}
			j++
		}
	}
	return res
}

func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
