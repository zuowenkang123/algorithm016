package Week_08

// https://leetcode-cn.com/problems/power-of-two/
// 2020-11-02 21:34:49

// 只存在一个1，通过减1与操作为0就可以判断
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
