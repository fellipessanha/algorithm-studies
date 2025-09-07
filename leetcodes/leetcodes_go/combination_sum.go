package leetcodesgo

func combinationSum(nums []int, target int) [][]int {
	answer := [][]int{}
	return backtrack(nums, target, []int{}, answer, 0)
}

func backtrack(nums []int, currentTarget int, path []int, answer [][]int, idx int) [][]int {
	if currentTarget == 0 {
		insertion := make([]int, len(path))
		copy(insertion, path)
		return append(answer, insertion)
	}

	if idx >= len(nums) || currentTarget < 0 {
		return answer
	}

	newTarget := currentTarget - nums[idx]
	answer = backtrack(nums, newTarget, append(path, nums[idx]), answer, idx)

	return backtrack(nums, currentTarget, path, answer, idx+1)
}
