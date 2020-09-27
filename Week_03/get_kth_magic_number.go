package Week_03

// https://leetcode-cn.com/problems/get-kth-magic-number-lcci/
// 2020-09-27

// 迭代的方式
func getKthMagicNumber(k int) int {
	if k <= 0 {
		return -1
	}
	magicNums := make([]int, k)
	magicNums[0] = 1
	p3, p5, p7 := 0, 0, 0
	for i := 1; i < k; i++ {
		// 比较三种情况大小，取最小的数据
		m := min(magicNums[p3]*3, magicNums[p5]*5)
		m = min(m, magicNums[p7]*7)
		magicNums[i] = m
		// 选择谁，谁加1
		if m == magicNums[p3]*3 {
			p3++
		}
		if m == magicNums[p5]*5 {
			p5++
		}
		if m == magicNums[p7]*7 {
			p7++
		}
	}
	return magicNums[k-1]
}
