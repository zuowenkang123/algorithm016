package Week_08

// https://leetcode-cn.com/problems/number-of-1-bits/
// 2020-11-02

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		// 清零最低位
		num = num & (num - 1)
	}
	return count
}

func hammingWeight1(num uint32) int {
	count := 0
	mask := uint32(1)
	for i := 0; i < 32; i++ {
		if (num & mask) != 0 {
			count++
		}
		mask <<= 1
	}
	return count
}
