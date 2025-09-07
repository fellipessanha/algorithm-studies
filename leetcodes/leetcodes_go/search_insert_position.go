package leetcodesgo

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	var l, r = 0, len(nums)
	for l <= r {
		m := l + (r-l)/2
		if m >= len(nums) || m <= 0 || nums[m] == target || (nums[m-1] < target && nums[m] > target) {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}
