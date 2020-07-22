/*Package leetcode 39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	var path []int
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	backtrack(candidates, 0, path, &ret, target)
	return ret
}

func backtrack(candidates []int, begin int, path []int, ans *[][]int, target int) {
	if target == 0 {
		n := make([]int, len(path))
		copy(n, path)
		*ans = append(*ans, n)
		return
	}
	for i := begin; i < len(candidates); i++ {
		residue := target - candidates[i]
		if residue < 0 {
			break
		}
		path = append(path, candidates[i])
		backtrack(candidates, i, path, ans, residue)
		path = path[:len(path)-1]
	}
	return
}
