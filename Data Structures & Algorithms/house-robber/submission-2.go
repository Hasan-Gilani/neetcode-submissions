func rob(nums []int) int {
    robn_2, robn_1 := 0, 0
	for i := 0; i < len(nums); i++ {
		currentLoot := nums[i]
		lootedSoFar := max(robn_2 + currentLoot, robn_1)
		robn_2 = robn_1
		robn_1 = lootedSoFar

	}
	return robn_1

}
