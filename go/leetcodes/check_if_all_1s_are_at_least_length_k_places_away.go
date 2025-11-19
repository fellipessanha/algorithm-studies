package leetcodesgo

func kLengthApart(nums []int, k int) bool {
	lastOne := -k - 1
	for idx, value := range nums {
		if value != 1 {
			continue
		}
		if idx-lastOne-1 < k {
			return false
		}
		lastOne = idx
	}
	return true
}
