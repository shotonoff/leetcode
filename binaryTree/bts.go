package binaryTree

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// 653. Two Sum IV - Input is a BST
func findTarget(root *TreeNode, k int) bool {
	memo := make(map[int]struct{})

	var result = false

	inOrder(root, func(root *TreeNode) {
		if _, ok := memo[k-root.Val]; ok {
			result = true
		}

		memo[root.Val] = struct{}{}
	})

	return result
}
