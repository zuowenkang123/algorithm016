package Week_08

// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
// 2020-11-02

func firstUniqChar(s string) int {
	charMap := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if charMap[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	arr := make([]int, 26)
	// 装入已有
	for i, k := range s {
		arr[k-'a'] = i
	}
	// 前面有会错乱该值，如果没有就不会错乱就对应上了
	for i, k := range s {
		if i == arr[k-'a'] {
			return i
		} else {
			arr[k-'a'] = -1
		}
	}
	return -1
}
