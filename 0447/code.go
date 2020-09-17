//447. 回旋镖的数量
//给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
//
//找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
//
//示例:
//
//输入:
//[[0,0],[1,0],[2,0]]
//
//输出:
//2
//
//解释:
//两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
package leetcode

func numberOfBoomerangs(points [][]int) int {
	var res int
	m := make(map[int]int)
	for i := 0; i < len(points); i++ {
		for k := range m {
			delete(m, k)
		}
		for j := 0; j < len(points); j++ {
			if j != i {
				m[dis(points[i], points[j])]++
			}
		}

		for _, v := range m {
			if v > 1 {
				res += v * (v - 1)
			}
		}
	}
	return res
}

func dis(a, b []int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return x*x + y*y
}
