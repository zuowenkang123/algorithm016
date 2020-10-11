package Week_06

// https://leetcode-cn.com/problems/house-robber/
// 2020-10-12
var robMap map[int]int

func rob(nums []int) int {
	robMap = make(map[int]int)
	return robD(nums, len(nums))
}

func robD(nums []int, n int) int {
	if _, ok := robMap[n]; ok {
		return robMap[n]
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	n2 := robD(nums, n-2) + nums[n-1]
	n1 := robD(nums, n-1)
	robMap[n] = max(n2, n1)
	return robMap[n]
}

func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// 空间复杂度O(1)
// 空间优化见：https://leetcode-cn.com/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	pre, res := 0, 0
	for _, i := range nums {
		res, pre = max(res, i+pre), res
	}
	return res
}
