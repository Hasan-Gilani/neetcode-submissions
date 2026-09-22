/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func deepClone(node *Node, visited map[*Node]*Node) *Node{
	if node == nil {
		return nil
	}
	if clone, ok := visited[node]; ok {
		return clone
	}
	clone := &Node{Val : node.Val}
	visited[node] = clone
	for _, n := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, deepClone(n, visited))
	}
	return clone

}

func cloneGraph(node *Node) *Node {
	cloned := make(map[*Node]*Node)
	return deepClone(node, cloned)

}
