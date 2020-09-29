package Week_04

// https://leetcode-cn.com/problems/lemonade-change/
// 2020-09-29
func lemonadeChange(bills []int) bool {
	billMap := make(map[int]int, 0)
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			billMap[5]++
		} else if bills[i] == 10 {
			if billMap[5] >= 1 {
				billMap[5]--
				billMap[10]++
			} else {
				return false
			}

		} else if bills[i] == 20 {
			if billMap[10] >= 1 && billMap[5] >= 1 {
				billMap[10]--
				billMap[5]--
				billMap[20]++
			} else if billMap[5] >= 3 {
				billMap[5] = billMap[5] - 3
				billMap[20]++
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
