package Week_02

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/group-anagrams/
// 2020-09-15

// 排序map法
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		charArr := []byte(str)
		charSlc := CharSlice(charArr)
		sort.Sort(charSlc)
		_, ok := strMap[string(charSlc)]
		if !ok {
			strMap[string(charSlc)] = []string{str}
		} else {
			strMap[string(charSlc)] = append(strMap[string(charSlc)], str)
		}
	}
	res := make([][]string, 0)
	for _, strArr := range strMap {
		res = append(res, strArr)
	}
	return res
}

// 字母累加法
func groupAnagrams1(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		keyArr := make([]byte, 26, 26)
		for _, v := range str {
			keyArr[v-'a']++
		}
		keyStr := ""
		for _, val := range keyArr {
			keyStr = fmt.Sprintf("%s%d#", keyStr, val)
		}
		if _, ok := strMap[keyStr]; !ok {
			strMap[keyStr] = []string{str}
		} else {
			strMap[keyStr] = append(strMap[keyStr], str)
		}
	}
	res := make([][]string, 0)
	for _, strArr := range strMap {
		res = append(res, strArr)
	}
	return res
}

// 质数法
func groupAnagrams2(strs []string) [][]string {
	return nil
}
