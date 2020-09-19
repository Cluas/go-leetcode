package leetcode

func maxPoints(points [][]int) int {
	// 点去重复 减少遍历次数 数组在golang可以作为map的key
	dm := make(map[[2]int]int)
	for _, point := range points {
		dm[[2]int{point[0], point[1]}]++
	}

	size := len(dm)

	// 如果不同的点个数小于等于2 直接返回点的数量
	if size <= 2 {
		return len(points)
	}

	// 构造不重复的坐标点
	diffPoints := make([][2]int, 0, size)
	for k := range dm {
		diffPoints = append(diffPoints, k)
	}

	// 三次遍历
	max := 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			count := 0
			for k := 0; k < size; k++ {
				if inLine(diffPoints[i], diffPoints[j], diffPoints[k]) {
					count += dm[diffPoints[k]]
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max

}

func inLine(x, y, z [2]int) bool {
	// 判断斜率 为了防止浮点数影响精度 改用乘法
	return (z[0]-x[0])*(y[1]-x[1]) == (y[0]-x[0])*(z[1]-x[1])
}
