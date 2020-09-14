package Week_01

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
// 2020-09-14

// map累加法
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
