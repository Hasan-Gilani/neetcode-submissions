//l = 1
//r = 4
// 2 


func minEatingSpeed(piles []int, h int) int {
	k := 0
	r := 0
	for _, p := range piles {
		r = max(r, p) 
	}
	l := 1
	for l <= r {
		eatingRate := l + (r - l) / 2
		totalHours := 0
		for _, p := range piles {
			totalHours += int(math.Ceil(float64(p)/float64(eatingRate)))
		}
		if totalHours > h {
			l = eatingRate + 1
		}
		if totalHours <= h {
			k = eatingRate
			r = eatingRate - 1
		}
	}
	return k
}
