## 几个小细节，几乎双百
### 解题思路
![image.png](https://pic.leetcode-cn.com/1600263691-VKsOTb-image.png)
总体思路是把four sum 转换成two sum。
其中有几个针对语言的优化点。
- 初始化字典内存，减少map由于growup内存导致的额外开销
- 巧妙利用 -a -b  之和等于 c d 之和的特性 （原题要求abcd相加为零所以一定存在这种组合）可以减少几次计算。


### 代码

```golang
func fourSumCount(A []int, B []int, C []int, D []int) int {
	ab := make(map[int]int, len(A) * len(A))
	for _, a := range A {
		for _, b:= range B {
			ab[-a-b]++
		}
	}
	var res int

	for _, c := range C {
		for _, d:= range D {
			if count, ok := ab[c+d]; ok {
				res += count
			}
		}
	}

	return res
}
```