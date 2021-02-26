package leetcode

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	// 前序遍历第一个节点是root
	// 从中序遍历中获取左右子树分割位置
	idx := findIdx(inorder, preorder[0])

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}

}

func findIdx(inorder []int, n int) int {
	for i, v := range inorder {
		if v == n {
			return i
		}
	}
	return 0
}
