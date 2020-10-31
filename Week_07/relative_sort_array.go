package Week_07

// https://leetcode-cn.com/problems/relative-sort-array/
// 2020-10-31

func relativeSortArray(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1), len(arr1))

	numMap := make(map[int]int, 0)
	for _, num := range arr1 {
		numMap[num]++
	}
	j := 0
	for _, data := range arr2 {
		for numMap[data] > 0 {
			res[j] = data
			j++
			numMap[data]--
		}
	}

	for i := 0; i < 1001; i++ {
		for numMap[i] > 0 {
			res[j] = i
			j++
			numMap[i]--
		}
	}
	return res
}
