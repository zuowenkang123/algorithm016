package Week_09

// https://leetcode-cn.com/problems/sort-an-array/
// 2020-12-01

//快速排序
func sortArray(nums []int) []int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	return nums
}

// 快排
func quickSort(nums []int, start int, end int) {
	//temirator
	if start > end {
		return
	}
	//processon
	pos := partition(nums, start, end)
	//drill down
	quickSort(nums, start, pos-1)
	quickSort(nums, pos+1, end)
}

func partition(nums []int, start int, end int) int {
	pivot := end

	pos := start
	//pivot作为标杆，counter左边都是小于pivo的值的
	for i := start; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
	nums[pos], nums[pivot] = nums[pivot], nums[pos]
	return pos
}

func mergeSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) >> 1
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	tmp := make([]int, end-start+1)
	i, j, k := start, mid+1, 0
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			k++
			i++
		} else {
			tmp[k] = nums[j]
			k++
			j++
		}
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= end {
		tmp[k] = nums[j]
		k++
		j++
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
}

func sortArray1(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}

	//set up the maxHeap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	//generate the order list from MaxHeap
	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
	return nums
}

func heapify(nums []int, n int, i int) {
	left, right := 2*i+1, 2*i+2
	largest := i

	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}
