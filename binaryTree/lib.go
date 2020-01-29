package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreeFromSlice(a []interface{}) *TreeNode {
	return helperBinaryTreeFromSlice(a, 0)
}

func helperBinaryTreeFromSlice(a []interface{}, i int) *TreeNode {
	if i >= len(a) || a[i] == nil {
		return nil
	}

	return &TreeNode{
		Val:   a[i].(int),
		Left:  helperBinaryTreeFromSlice(a, 2*i+1),
		Right: helperBinaryTreeFromSlice(a, 2*i+2),
	}
}

func inOrder(root *TreeNode, visitor func(root *TreeNode)) {
	var st []*TreeNode

	for root != nil || len(st) > 0 {
		if root == nil {
			var l = len(st) - 1
			root = st[l]
			visitor(root)
			root = root.Right
			st = st[:l]

		} else {
			st = append(st, root)
			root = root.Left
		}
	}
}
