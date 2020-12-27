package Week_06

// https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	rowNum, colNum := len(matrix), len(matrix[0])
	result := math.MinInt32
	// 枚举左边界
	for left := 0; left < colNum; left++ {
		// 左边界改变才算区域的重新开始
		rowSum := make([]int, rowNum)
		for right := left; right < colNum; right++ {
			// 将二维数组降为一维数组处理
			for row := 0; row < rowNum; row++ {
				rowSum[row] += matrix[row][right]
			}
			result = max(result, maxSubArrBelowK(rowSum, k))
			if result == k {
				return k
			}
		}
	}
	return result
}

// 在数组 arr 中，求不超过 k 的最大值
func maxSubArrBelowK(arr []int, k int) int {
	sum, max, l := arr[0], arr[0], len(arr)
	// O(rows)
	for i := 1; i < l; i++ {
		if sum > 0 {
			sum += arr[i]
		} else {
			sum = arr[i]
		}
		if sum > max {
			max = sum
		}
	}
	if max <= k {
		return max
	}
	// O(rows ^ 2)
	max = math.MinInt32
	for left := 0; left < l; left++ {
		sum := 0
		for right := left; right < l; right++ {
			sum += arr[right]
			if sum > max && sum <= k {
				max = sum
			}
			if max == k {
				return k
			}
		}
	}
	return max
}
