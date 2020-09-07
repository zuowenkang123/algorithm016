package Week_01

var stairsMap map[int]int

// 递归解法
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

// 动态规划,递推逻辑
func climbStairsDp(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	stairArr := make([]int, 0, n)
	// 初始数据
	stairArr = append(stairArr, 1)
	stairArr = append(stairArr, 2)
	// 不存在则赋值
	for i := 2; i < n; i++ {
		stairArr = append(stairArr, stairArr[i-1]+stairArr[i-2])
	}

	return stairArr[n-1]
}

// 动态规划,递推逻辑,空间复杂度下降
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
