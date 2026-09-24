type KthLargest struct {
    k int
	nums []int
}


func Constructor(k int, nums []int) KthLargest {
    kl := KthLargest{
		k: k,
		nums: []int{},
	}
	
	for _, v := range(nums){
		kl.Add(v)
	}
	return kl
}


func (this *KthLargest) Add(val int) int {
    this.nums = append(this.nums, val)
	this.siftUp(len(this.nums) - 1)

	if len(this.nums) > this.k {
		this.nums[0] = this.nums[len(this.nums)-1]
		this.nums = this.nums[:len(this.nums)-1]
		i := 0
		for {
			l := 2 * i + 1
			r := 2 * i + 2
			smallest := i
			if l < len(this.nums) && this.nums[l] < this.nums[smallest] {
				smallest = l
			}
			if r < len(this.nums) && this.nums[r] < this.nums[smallest] {
				smallest = r
			}
			if smallest == i {
				break
			}
			this.nums[i], this.nums[smallest] = this.nums[smallest], this.nums[i]
			i = smallest
		}
	}
	return this.nums[0]
}

func (this *KthLargest) siftUp(i int) {
	for i > 0 && this.nums[(i-1)/2] > this.nums[i] {
		parent := (i-1)/2
		this.nums[parent], this.nums[i] = this.nums[i], this.nums[parent]
		i = parent
	}
}