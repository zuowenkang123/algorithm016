package Week_03

import (
	"math/rand"
	"sort"
)

// https://leetcode-cn.com/problems/majority-element/description/
// 2020-09-24

// map法
func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for num, val := range numMap {
		if val > len(nums)/2 {
			return num
		}
	}
	return 0
}

// 排序
func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 投票
func majorityElement2(nums []int) int {
	mode := nums[0]
	votes := 0
	for _, num := range nums {
		if votes == 0 {
			mode = num
		}
		if num == mode {
			votes++
		} else {
			votes--
		}
	}
	return mode
}

// 随机化
func majorityElement3(nums []int) int {

	halfLen := len(nums) / 2
	for {
		randIdx := randomFunc(0, len(nums))
		num := nums[randIdx]
		if findResult(nums, num) > halfLen {
			return num
		}
	}
}

func randomFunc(min, max int) int {
	return rand.Intn(max-min) + min
}

func findResult(nums []int, num int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}
