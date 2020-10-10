package Week_06

// https://leetcode-cn.com/problems/search-a-2d-matrix/
// 2020-10-10

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	left, right := 0, m*n-1
	mid, midValue := 0, 0
	for left <= right {
		mid = (left + right) / 2
		midValue = matrix[mid/n][mid%n]
		if target == midValue {
			return true
		} else if target < midValue {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false

}
