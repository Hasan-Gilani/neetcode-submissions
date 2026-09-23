func robNonAdjacent(nums []int) int {
	robn_2, robn_1 := 0, 0
	for _, n := range nums {
		currentLoot := max(robn_2 + n, robn_1)
		robn_2 = robn_1
		robn_1 = currentLoot
	}
	return robn_1
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
    n := len(nums)
	return max(robNonAdjacent(nums[:n-1]), robNonAdjacent(nums[1:n]))
	
	
}
