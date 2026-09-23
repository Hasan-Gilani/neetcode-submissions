/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    output := []int{}
	queue := []*TreeNode{root}
	levelValue := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			x := queue[0]
			queue = queue[1:]
			levelValue = x.Val
			if x.Left != nil {
				queue = append(queue, x.Left)
			} 
			if x.Right != nil {
				queue = append(queue, x.Right)
			}
		}
		output = append(output, levelValue)
	}
	return output
}
