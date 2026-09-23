/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
    output := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n ; i ++ {
			a := queue[0]
			queue = queue[1:]
			level = append(level, a.Val)
			if a.Left != nil {
				queue = append(queue, a.Left)
			}
			if a.Right != nil {
				queue = append(queue, a.Right)
			}
		}
		output = append(output, level)
	}
	return output
}
