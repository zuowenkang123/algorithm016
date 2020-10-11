package Week_06

// https://leetcode-cn.com/problems/house-robber-ii/
// 2020-10-12
// https://leetcode-cn.com/problems/house-robber-ii/solution/tong-yong-si-lu-tuan-mie-da-jia-jie-she-wen-ti-by-/

func rob_2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(myRob(nums[1:]), myRob(nums[:len(nums)-1]))
}

func myRob(nums []int) int {
	cur, pre := 0, 0
	for _, v := range nums {
		cur, pre = max(pre+v, cur), cur
	}
	return cur
}
