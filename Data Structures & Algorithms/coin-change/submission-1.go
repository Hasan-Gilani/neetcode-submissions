func coinChange(coins []int, amount int) int {
	solutions := make([]int, amount + 1)
	solutions[0] = 0
	for i := 1; i < len(solutions); i++ {
		solutions[i] = 10001
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i >= c {
				solutions[i] = min(1 + solutions[i-c], solutions[i])
			}
		}
	}
	lastElement := solutions[len(solutions)-1]
	if lastElement > amount {
		return -1
	} else {
		return lastElement
	}
}
