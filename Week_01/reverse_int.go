package Week_01

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/reverse-integer/
// 2020-10-14

func reverseInt(x int) int {
	flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	arrInt := make([]int, 0)
	for x > 0 {
		arrInt = append(arrInt, x%10)
		x = x / 10
	}
	res := 0
	for _, i := range arrInt {
		res = res*10 + i
	}
	res = res * flag
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}

func reverseInt2(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res < math.MinInt32 || res > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}

	return res
}

// 小朱的问题
func reverseInt1(arrList [][]int) map[int][]int {
	// 根相同的数组
	arrMap := make(map[int][]int, 0)
	// 填入首个根数组
	for _, arr := range arrList {
		if _, ok := arrMap[arr[0]]; !ok {
			arrMap[arr[0]] = arr
		}
	}
	for _, arr1 := range arrList {
		arrRoot := arrMap[arr1[0]]
		for i := 0; i < len(arr1); i++ {
			// 新数组存在0而根数组不为0，则替换
			if arr1[i] == 0 && arrRoot[i] != 0 {
				arrMap[arr1[0]] = arr1
				break
			}
		}
	}
	fmt.Println(arrMap)
	return arrMap
}
