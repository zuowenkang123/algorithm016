package Week_01

// https://leetcode-cn.com/problems/intersection-of-two-arrays/submissions/
// 2020-12-04
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	nums1Map := make(map[int]int)
	for _, val := range nums1 {
		nums1Map[val] = val
	}
	nums2Map := make(map[int]int)
	for _, val := range nums2 {
		nums2Map[val] = val
	}
	res := make([]int, 0)
	for _, val := range nums1Map {
		if _, ok := nums2Map[val]; ok {
			res = append(res, val)
		}
	}
	return res
}
