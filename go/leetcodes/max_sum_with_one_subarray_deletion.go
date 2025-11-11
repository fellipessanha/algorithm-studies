package leetcodesgo

func maximumSum(arr []int) int {
	maxWithErases := arr[0]
	maxWithoutErases := arr[0]
	maxSum := arr[0]

	for _, n := range arr[1:] {
		maxWithErases = max(maxWithoutErases, maxWithErases+n)
		maxWithoutErases = max(maxWithoutErases+n, n)
		maxSum = max(maxSum, maxWithErases, maxWithoutErases)
	}

	return maxSum
}
