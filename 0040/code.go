/*Package code 40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

*/
package code

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// 排序，以减少搜索的分支
	sort.Ints(candidates)
	// dfs 搜索
	return dfs(candidates, target)
}

func dfs(nums []int, target int) [][]int {
	var ans [][]int
	for i, n := range nums {
		// 与上一个值比较是否相同，避免重复搜索
		if i > 0 && nums[i-1] == n {
			continue
		} else if target < n {
			break
		} else if target == n {
			ans = append(ans, []int{n})
			continue
		}
		// 不能重复使用数字组合
		for _, v := range dfs(nums[i+1:], target-n) {
			ans = append(ans, append([]int{n}, v...))
		}
	}
	return ans
}
