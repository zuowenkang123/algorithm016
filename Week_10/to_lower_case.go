package Week_10

// https://leetcode-cn.com/problems/to-lower-case/solution/san-chong-fang-fa-by-lyai-bian-cheng/

func toLowerCase(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] = bytes[i] + 32
		}
	}
	return string(bytes)
}
