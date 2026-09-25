func longestConsecutive(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	numbers := make(map[int]struct{})
	maxSequence := 0 
	for _, n := range nums {
		numbers[n] = struct{}{}
	}
	for i, _ := range numbers {
		if _, ok := numbers[i-1]; !ok{
			sequence := 1
			j := i
			for {
				j++
				if _, ok := numbers[j]; !ok {
					break
				}
				sequence++
			}
			maxSequence = max(sequence, maxSequence)
		}
		
	}
	return maxSequence
}
