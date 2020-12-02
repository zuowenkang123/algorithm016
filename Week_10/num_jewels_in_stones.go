package Week_10

// https://leetcode-cn.com/problems/jewels-and-stones/
// 2020-12-02

func numJewelsInStones(J string, S string) int {
	jMap := make(map[int32]bool, 0)
	for _, v := range J {
		jMap[v] = true
	}
	res := 0
	for _, s := range S {
		if _, ok := jMap[s]; ok {
			res++
		}
	}
	return res
}
