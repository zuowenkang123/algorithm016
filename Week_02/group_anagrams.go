package Week_02

import "sort"

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
			strArr := make([]string, 0)
			strArr = append(strArr, str)
			strMap[string(charSlc)] = strArr
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
