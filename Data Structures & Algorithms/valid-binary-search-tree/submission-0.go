/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var validate func(node *TreeNode, min, max *int) bool
	validate = func(node *TreeNode, min, max *int) bool {
		if node == nil {
			return true
		}
		if min != nil && node.Val <= *min {
			return false
		}
		if max != nil && node.Val >= *max {
			return false
		}
		return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
	}
	return validate(root, nil, nil)
}
