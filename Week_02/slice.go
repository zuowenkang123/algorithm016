package Week_02

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
