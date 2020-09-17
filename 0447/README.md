## 447. 回旋镖的数量

### 解题思路
利用hash表把O(n^3)降低到O(n^2)
有个节省内存的小细节，只初始化一次记录表，每次循环开始前清空map

### 代码

```go
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
```