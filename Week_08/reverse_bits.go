package Week_08

// https://leetcode-cn.com/problems/reverse-bits/
// 2020-11-02

func reverseBits(num uint32) uint32 {
	power, ret := uint32(31), uint32(0)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power--
	}
	return ret
}
