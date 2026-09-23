type Node struct {
	isEnd bool
	child [26]*Node
}

type PrefixTree struct {
	root *Node
}

func Constructor() PrefixTree {
    return PrefixTree{root: &Node{}}
}

func (this *PrefixTree) Insert(word string) {
	node := this.root
	for _, w := range word {
		idx := w - 'a'
		if node.child[idx] == nil {
			node.child[idx] = &Node{}
		}
	node = node.child[idx]
	}
	node.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this.root
	for _, w := range word {
		idx := w - 'a'
		if node.child[idx] == nil {
			return false
		}
		node = node.child[idx]
	}
	return node.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.root
	for _, w := range prefix {
		idx := w - 'a'
		if node.child[idx] == nil {
			return false
		}
		node = node.child[idx]
	}
	return true
}
