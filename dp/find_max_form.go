package dp

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		countArr := GetCount4ZeroAndOne(str)
		for zero := m; zero >= countArr[0]; zero-- {
			for one := n; one >= countArr[1]; one-- {
				dp[zero][one] = max(dp[zero-countArr[0]][one-countArr[1]]+1, dp[zero][one])
			}
		}
	}
	return dp[m][n]
}

func GetCount4ZeroAndOne(s string) map[byte]int {
	countArr := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		countArr[s[i]-'0']++
	}
	return countArr
}
