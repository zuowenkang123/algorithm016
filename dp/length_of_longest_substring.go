package dp

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
// 2020-12-22 21:03:45

func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	left, right := 0, 0
	length := 0
	for right < len(s) {
		if _, ok := window[s[right]]; ok {
			//发现了重复
			delete(window, s[left])
			left++
		} else {
			if right-left >= length {
				length = right - left + 1
			}
			//不存在重复
			window[s[right]] = 1
			right++
		}
	}
	return length
}

/*
 * DP:时间复杂度O(n)，空间复杂度O(1).
 * dp[i]表示字符s[i]结尾的最长不重复子串长度。
 * index表示s[i]左边距离最近的相同字符。
 */
func lengthOfLongestSubstring1(s string) int {
	strMap := make(map[byte]int)
	res := 0
	dp := 0
	for i := 0; i < len(s); i++ {
		dupIndex, ok := strMap[s[i]]
		if ok == false {
			dupIndex = -1
		}
		strMap[s[i]] = i
		// 前面没有重复
		if dp < i-dupIndex {
			dp += 1
		} else {
			// 当前相同字母当上一个字母的长度
			dp = i - dupIndex
		}
		res = max(res, dp)
	}
	return res
}
