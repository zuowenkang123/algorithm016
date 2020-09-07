package Week_01

var stairsMap map[int]int

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
