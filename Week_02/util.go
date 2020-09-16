package Week_02

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 字符排序用
type CharSlice []byte

func (this CharSlice) Len() int {
	return len(this)
}

func (this CharSlice) Less(i, j int) bool {
	if this[i] > this[j] {
		return false
	}
	return true
}

func (this CharSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	return
}
