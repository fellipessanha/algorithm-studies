package leetcodesgo

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    answers := make([][]int, 0)
    for i := range nums{
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }

        // two pointer strategy for sums!
        l, r := i+1, len(nums)-1
        for l < r{
            if nums[i] + nums[l] + nums[r] == 0 {
                answers = append(answers, []int{nums[i], nums[l], nums[r]})
                l++
            } else if nums[i] + nums[l] + nums[r] > 0 {
                r--
            } else{
                l++
            }
        }
    }
    return answers
}

