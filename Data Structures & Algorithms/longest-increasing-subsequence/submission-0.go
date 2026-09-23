func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums) + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	m := -1
	for _, v := range dp {
		m = max(m, v)
	}
	return m
}
