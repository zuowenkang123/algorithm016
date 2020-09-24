package Week_03

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 2020-09-24

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string
var combination string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	combination = ""
	backtracketterCombinations(digits, 0)
	return combinations
}

func backtracketterCombinations(digits string, index int) {
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	}
	digit := string(digits[index])
	letters := phoneMap[digit]
	for i := 0; i < len(letters); i++ {
		combination += string(letters[i])
		backtracketterCombinations(digits, index+1)
	}
}
