type KthLargest struct {
    k int
	h []int
}


func Constructor(k int, nums []int) KthLargest {
    kl := KthLargest{
		k: k,
		h : []int{},
	}
	for _, n := range nums{
		kl.Add(n)
	}
	return kl
}

func (this KthLargest) Less(i, j int) bool{
	return this.h[i] < this.h[j]
}

func (this *KthLargest) Swap(i, j int){
	this.h[i], this.h[j] = this.h[j], this.h[i]
}

func (this *KthLargest) Pop() {
	l := len(this.h)
	this.h[0] = this.h[l-1]
	this.h = this.h[:l-1]
}

func (this *KthLargest) SiftUp(i int) {
	for i > 0 && this.Less(i, (i-1)/2) {
		parent := (i-1)/2
		this.Swap(i, parent)
		i = parent
	}
}

func (this *KthLargest) SiftDown(i int) {
	for {
		l, r, smallest := 2*i + 1, 2 * i + 2, i
		if l < len(this.h) && this.Less(l, smallest){
			smallest = l
		}
		if r < len(this.h) && this.Less(r, smallest){
			smallest = r
		}
		if smallest == i {
			break
		}
		this.Swap(smallest, i)
		i = smallest
		
	}
}

func (this *KthLargest) Add(val int) int {
	this.h = append(this.h, val)
	l := len(this.h)
	this.SiftUp(l-1)
	if l > this.k {
		this.Pop()
		this.SiftDown(0)
	}

	return this.h[0]
}
