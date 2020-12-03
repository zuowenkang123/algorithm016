package Week_10

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/word-pattern/
// 2020-12-03
func wordPattern(pattern string, str string) bool {
	p := strings.Split(pattern, "")
	s := strings.Split(str, " ")
	if len(p) != len(s) {
		return false
	}
	pNum, sNum := 0, 0
	pString, sString := "", ""
	pMap := map[string]int{}
	sMap := map[string]int{}
	for _, v := range p {
		if _, ok := pMap[v]; ok {
			pString += strconv.Itoa(pMap[v])
		} else {
			pString += strconv.Itoa(pNum)
			pMap[v] = pNum
			pNum++
		}
	}
	for _, v := range s {
		if _, ok := sMap[v]; ok {
			sString += strconv.Itoa(sMap[v])
		} else {
			sString += strconv.Itoa(sNum)
			sMap[v] = sNum
			sNum++
		}
	}
	if pString == sString {
		return true
	}
	return false
}
