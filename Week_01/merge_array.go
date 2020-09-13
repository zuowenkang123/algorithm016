package Week_01

// https://leetcode-cn.com/problems/merge-sorted-array/
// 2020-09-13

// 安排法，每个位置肯定需要有人坐。
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for pos := m + n - 1; pos >= 0; pos-- {
		if i < 0 {
			nums1[pos] = nums2[j]
			j--
		} else if j < 0 {
			nums1[pos] = nums1[i]
			i--
		} else if nums1[i] >= nums2[j] {
			nums1[pos] = nums1[i]
			i--
		} else {
			nums1[pos] = nums2[j]
			j--
		}
	}
}

// 放置加和法
func merge1(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1] = nums1[i]
			i--
		}
	}

	//若剩下nums1，直接就是其中的值；若剩下nums2，则倒序依次填进去
	for j >= 0 {
		nums1[j] = nums2[j]
		j--
	}
}

// 小数组循环，主被动关系
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		// 当nums1循环为0时，直接填入nums2。两种情况，一种是一开始为0，另一种是移动后为0.
		if m == 0 {
			nums1[n-1] = nums2[n-1]
			n -= 1
			continue
		}
		if nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n -= 1
		} else {
			nums1[m+n-1] = nums1[m-1]
			m -= 1
		}
	}
}

// 同上，判断逻辑确认
func merge3(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for j >= 0 {
		if i < 0 {
			nums1[j] = nums2[j]
			j--
			continue
		}
		if nums1[i] < nums2[j] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1] = nums1[i]
			i--
		}
	}
}
