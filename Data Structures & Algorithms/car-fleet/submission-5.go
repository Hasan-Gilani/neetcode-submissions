//4 -> ahead
//1 -> behind
//4, 1
func carFleet(target int, position []int, speed []int) int {
	indexes := []int{}
	for i := range position {
		indexes = append(indexes, i)
	}
	sort.Slice(indexes, func(i, j int) bool {
		return position[indexes[i]] > position[indexes[j]]
	})
	
	arrivals := []float64{}
	for _, i := range indexes {
		arrival := float64(target - position[i]) / float64(speed[i])
		if len(arrivals) == 0 || arrival > arrivals[len(arrivals) - 1] {
			arrivals = append(arrivals, arrival)
		}
	}
	return len(arrivals)
}
