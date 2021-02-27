package leetcode

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// 保存全局最大值，注意不能为零因为节点可能为复数
	maxSum := math.MinInt32
	var _maxPathSum func(root *TreeNode) int
	_maxPathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 注意最大值为负数的时候不加才是最优
		left := max(0, _maxPathSum(root.Left))
		right := max(0, _maxPathSum(root.Right))
		// 节点左中右和的值
		// 如果比最大值要大 更新这个最大值
		maxSum = max(root.Val+left+right, maxSum)
		// 取左右路径中较大的值
		return root.Val + max(left, right)
	}
	_maxPathSum(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
