package Week_04

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
// 2020-10-11

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		// 判断是否在前半部分查找
		if (nums[low] <= target && target <= nums[mid]) || (nums[mid] <= nums[high] && (target < nums[mid] || target > nums[high])) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// [left,mid] 连续递增
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1 // 在左侧 [left,mid) 查找
			} else {
				left = mid + 1
			}
		} else { // [mid,right] 连续递增
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1 // 在右侧 (mid,right] 查找
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
