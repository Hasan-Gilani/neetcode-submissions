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
    answer := []int{}
	queue := []*TreeNode{root}
	answer = append(answer, root.Val)
	for len(queue) != 0 {
		nodes := len(queue)
		ans := -101
		for nodes > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
					ans = node.Left.Val
					queue = append(queue, node.Left)
				}
			if node.Right != nil {
					ans = node.Right.Val
					queue = append(queue, node.Right)
				}
			nodes--
		}
		if ans > -101 {
			answer = append(answer, ans)
		}
	}
	return answer
}
