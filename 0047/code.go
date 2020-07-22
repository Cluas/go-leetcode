/*Package leetcode 47. 全排列 II
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/
package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	var mark = make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums, path, mark, &res)
	return res
}

func backtrack(nums []int, path []int, mark []bool, res *[][]int) {
	if len(nums) == len(path) {
		n := make([]int, len(nums))
		copy(n, path)
		*res = append(*res, n)
		return
	}
	for i := 0; i < len(nums); i++ {
		if mark[i] || (i > 0 && nums[i] == nums[i-1] && !mark[i-1]) {
			continue
		}
		mark[i] = true
		path = append(path, nums[i])
		backtrack(nums, path, mark, res)
		mark[i] = false
		path = path[:len(path)-1]

	}
	return
}
