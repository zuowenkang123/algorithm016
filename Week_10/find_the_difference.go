package Week_10

// https://leetcode-cn.com/problems/find-the-difference/
// 2020-12-02
func findTheDifference(s string, t string) byte {
	ssum, tsum := 0, 0
	for _, x := range s {
		ssum += int(x)
	}
	for _, x := range t {
		tsum += int(x)
	}
	return byte(tsum - ssum)
}
