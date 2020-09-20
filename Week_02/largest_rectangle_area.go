package Week_02

import "math"

//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
// 2020-09-20

// 栈
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	stack := StackInt{}
	// 需要对比用
	stack.Push(-1)

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for stack[len(stack)-1] != -1 && heights[stack.Top()] >= heights[i] {
			cur := stack.Top()
			stack.Pop()
			maxArea = max(maxArea, heights[cur]*(i-stack.Top()-1))
		}
		stack.Push(i)
	}

	for stack[len(stack)-1] != -1 {
		cur := stack.Top()
		stack.Pop()
		maxArea = max(maxArea, heights[cur]*(len(heights)-stack.Top()-1))
	}
	return maxArea
}

func largestRectangleArea2(heights []int) int {

	l := len(heights)

	if l == 0 {
		return 0
	}
	if l == 1 {
		return heights[0]
	}

	// 用切片模拟栈，为了节省切片扩容时间，预设切片空间为数组长度
	// 栈用来存heights数组中元素下标
	stack := make([]int, 1, l)
	stack[0] = -1 // 预填一个-1

	maxArea := 0
	var area int
	for i := 0; i < l; i++ {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			area = heights[stack[len(stack)-1]] * (i - stack[len(stack)-2] - 1)
			if area > maxArea {
				maxArea = area
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 { // 栈未到底
		area = heights[stack[len(stack)-1]] * (l - stack[len(stack)-2] - 1)
		if area > maxArea {
			maxArea = area
		}
		stack = stack[:len(stack)-1]
	}
	return maxArea
}

func largestRectangleArea3(heights []int) int {
	// 1.暴力（根据宽找高度）：
	// 两个柱子之间的最大矩形，由最矮柱子决定，所以在两个指针的遍历中找到最矮柱子
	// 2.分治：最矮柱子的矩形宽可以无限延伸。找到左右边界范围内的最矮柱子，求以它为高的矩形、其左、右柱子中（子问题）矩形最大面积的最大值。
	// 3.优化分治：用线段树代替遍历找到最矮柱子。
	// 4.升序栈（根据高度找宽）：对于每个高度来说，往左右两个方向寻找第一个比它矮的柱子，决定矩形的宽。
	// 所以用高度升序栈保存索引，遇到减小的，表示找到右边界。计算这中间每个柱子的最大矩形，计算一个即出栈，这样方便每次从栈顶取索引。
	// 时间复杂度：O(n)， n个数字每个会被压栈弹栈各一次。
	// 空间复杂度：O(n)。用来存放栈中元素。

	// 首尾添加负数高度，这样原本的第一个高度能形成升序，原本的最后一个高度也能得到处理
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)
	size := len(heights)
	// 递增栈
	s := make([]int, 1, size)

	res := 0
	i := 1
	for i < len(heights) {
		// 递增则入栈
		if heights[s[len(s)-1]] < heights[i] {
			s = append(s, i)
			i++
			continue
		}
		// s[len(s)-2]是矩形的左边界
		res = max(res, heights[s[len(s)-1]]*(i-s[len(s)-2]-1))
		s = s[:len(s)-1]
	}
	return res
}

// 柱状图中最大的矩形
// n个非负整数，表示柱状图中柱高，每个柱子宽度为1
// 求能勾勒出的矩形的最大面积

// 这题和装水容器、接雨水比较像，都是左右双指针内移的套路

// 根据题意
// 矩形有两类情况
// 1. 取heights最小值，以heights长度为宽，所得矩形面积为area1
// 2.
// 要取最大矩形面积 s=b(宽)*h(高) ，则肯定先取其中一边 b 为max，再将 b 慢慢缩小，h则尝试向增大的方向变化，才可能找到最大矩形面积
// 这里h是宽度范围里的最小值，当b缩小时，需要更新h值

// 当区间内包含 minHeight 时， h 必定为 min， 再怎么移动左右指针都不可能增大矩形面积， 所以每次都要寻找最小值
// 假设一开始heights有最小值nums[min]，其下标min，此时矩形值为 s0 = len(heights) * nums[min]
// 计算左半区间(不含min)的新最小值min1，并计算新的矩形值s1 = min * min1，
// 如果 s1 <= s0， 则左边，重复前边的操作：按min1分成新的两个区间，再去计算
// 如果 s1 > s0， 更新 maxS = s1, 并在 min1 左右分成两个区间看能不能找到更大的面积
// 从这个分析来看，适合递归求解

// 1. 自己给出的递归分治解法
// 不考虑内层寻找最小值的情况的话。外层其实类似二分查找，时间复杂度应该是O(logn)，但恶化的情况下会变成O(n)
// 而每一次findMaxArea都要O(n)时间。
// 时间O(nlogn),空间O(n)
//96/96 cases passed (664 ms)
//Your runtime beats 31.25 % of golang submissions
//Your memory usage beats 7.14 % of golang submissions (7.2 MB)
func largestRectangleArea1(heights []int) int {

	// 1. 特殊情况
	l := len(heights)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return heights[0]
	}

	// 2. 一般情况
	var maxArea int
	findMaxArea(heights, &maxArea)
	return maxArea
}

func findMaxArea(ints []int, maxArea *int) {
	l := len(ints)
	if l == 0 {
		return
	} // 当ints没有数之时返回
	minI, minV := minInInts(ints)
	area := l * minV
	if area > *maxArea {
		*maxArea = area
	} // 更新最大面积
	// 将区间一分为二再计算
	if minI != 0 || minI != l-1 {
		findMaxArea(ints[:minI], maxArea)
		findMaxArea(ints[minI+1:], maxArea)
	} else if minI == 0 {
		findMaxArea(ints[1:], maxArea)
	} else if minI == len(ints)-1 {
		findMaxArea(ints[:l-1], maxArea)
	}

}

// 这里是采用一次扫描的办法取最小值，但这样显然浪费了许多计算，先这样写着
func minInInts(ints []int) (index, value int) {
	value = math.MaxInt32
	for i := 0; i < len(ints); i++ {
		if ints[i] < value {
			index, value = i, ints[i]
		}
	}
	return index, value
}

// 3. 暴力解法
// 遍历整个heights，当遍历到heights[i]时，看包含heights[i]在内的连续的>=heights[i]的柱子有多少
// 若有 n 个，则矩形宽即为 n ， 则 S = heights[i] * n
// 如果每次都去找的话需要 O(n^2)时间， O(1)空间
// 如果再对heights作去重，使用哈希集合记录不重复高度，则空间占用O(n)
//96/96 cases passed (248 ms)
//Your runtime beats 46.64 % of golang submissions
//Your memory usage beats 50 % of golang submissions (5.3 MB)
func largestRectangleArea4(heights []int) int {
	var result int
	l := len(heights)
	for k, v := range heights {
		left, right := k, k
		for left >= 0 && heights[left] >= v {
			left--
		}
		for right < l && heights[right] >= v {
			right++
		}
		result = max(result, (right-left-1)*v)
	}
	return result
}
