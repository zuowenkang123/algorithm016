package Week_02

// https://leetcode-cn.com/problems/valid-anagram/description/
// 2020-09-15

// map注意主被关系
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	strMap := make(map[int32]int)
	for _, v := range s {
		strMap[v]++
	}
	for _, v := range t {
		val, ok := strMap[v]
		if !ok || val == 0 {
			return false
		}
		strMap[v]--
	}
	return true
}

// 抵消法
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countMap := make(map[byte]int)
	for i, _ := range s {
		countMap[s[i]]++
		countMap[t[i]]--
	}
	for _, count := range countMap {
		if count != 0 {
			return false
		}
	}
	return true
}

// 排序字符串，等位对比
func isAnagram2(s string, t string) bool {
	return true
}
