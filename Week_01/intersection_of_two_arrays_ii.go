package Week_01

import "sort"

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
// 2020-09-14

// map累加法
// 时间 O(max(n,m)) 空间 O(min(n,m))
func intersect(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		if _, ok := numMap[nums1[i]]; ok {
			numMap[nums1[i]]++
		} else {
			numMap[nums1[i]] = 1
		}
	}
	intersectArr := make([]int, 0, len(nums1))
	for j := 0; j < len(nums2); j++ {
		_, ok := numMap[nums2[j]]
		if ok && numMap[nums2[j]] != 0 {
			intersectArr = append(intersectArr, nums2[j])
			numMap[nums2[j]]--
		}
	}
	return intersectArr
}

// 排序双指针
// 时间O(max(nlogn,mlogm,n+m)) 空间O(1)
func intersect1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	pos1 := 0
	pos2 := 0
	intersectArr := make([]int, 0, len(nums1))
	for pos1 < len(nums1) && pos2 < len(nums2) {
		if nums1[pos1] == nums2[pos2] {
			intersectArr = append(intersectArr, nums1[pos1])
			pos1++
			pos2++
		} else if nums1[pos1] > nums2[pos2] {
			pos2++
		} else {
			pos1++
		}
	}
	return intersectArr
}
