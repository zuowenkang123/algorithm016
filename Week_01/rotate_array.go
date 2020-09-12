package Week_01

// https://leetcode-cn.com/problems/rotate-array/
// 2020-09-12

// 旋转数组,循环处理 O(n*k) O(1)
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}

// 反转方法
// 反转全部，反转前k个，反转后面
func rotate1(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 环形替换
func rotate2(nums []int, k int) {
	k %= len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		from := start
		fromItem := nums[start]
		for count < len(nums) {
			to := (from + k) % len(nums)
			tmp := nums[to]
			nums[to] = fromItem
			fromItem = tmp
			from = to
			count++
			if start == from {
				break
			}
		}
	}
}
