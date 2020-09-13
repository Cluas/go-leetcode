![截屏2020-09-13 下午8.49.40.png](https://pic.leetcode-cn.com/1600001404-ovpYHv-%E6%88%AA%E5%B1%8F2020-09-13%20%E4%B8%8B%E5%8D%888.49.40.png)
### 解题思路
首先我们把后面的字符串按照空格分割，再做一些简单的特殊情况判定。
然后分别构建两个map，一个存储**pattern**每个对应位置的字符到**ss**对应字符串的映射，另一个相反操作

开始遍历**pattern**，分别判定查找表里面的值等不等于对应映射，遇到不相等直接返回false。
遍历结束后，说明所有映射符合模式，返回true

### 题解
```go
func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}

	p2s := make(map[byte]string)
	s2p := make(map[string]byte)

	for i := 0; i < len(pattern); i ++ {
		if v, ok := p2s[pattern[i]]; ok && v != ss[i] {
			return false
		}
		if v, ok := s2p[ss[i]]; ok && v != pattern[i] {
			return false
		}
		p2s[pattern[i]] = ss[i]
		s2p[ss[i]] = pattern[i]
	}
	return true
}
```