package dp

// https://leetcode-cn.com/problems/climbing-stairs/
// 2020-09-07
var stairsMap map[int]int

// 递归解法，记忆化搜索法
func climbStairs(n int) int {
	// 初始数据
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	// 初始化map
	if stairsMap == nil {
		stairsMap = make(map[int]int)
	}
	// 不存在则赋值
	if _, ok := stairsMap[n]; !ok {
		stairsMap[n] = climbStairs(n-1) + climbStairs(n-2)
	}

	return stairsMap[n]
}

// 动态规划,递推解法，数组填入法
func climbStairsDp(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n)
	// 初始数据
	dp[0] = 1
	dp[1] = 2
	// 不存在则赋值
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// 动态规划,递推解法,前后移动法
func climbStairsDp2(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	pre := 1
	suf := 2
	result := 0
	for i := 3; i <= n; i++ {
		result = pre + suf
		pre, suf = suf, result
	}

	return result
}

// 动态规划,递推解法，数组填入法
func climbStairsDp3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	// 初始数据
	dp[1] = 1
	dp[2] = 2
	// 不存在则赋值
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
