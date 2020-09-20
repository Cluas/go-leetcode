### 217. 存在重复元素
<div class="notranslate"><p>给定一个整数数组，判断是否存在重复元素。</p>

<p>如果任意一值在数组中出现至少两次，函数返回 <code>true</code> 。如果数组中每个元素都不相同，则返回 <code>false</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [1,2,3,1]
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>[1,2,3,4]
<strong>输出:</strong> false</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入: </strong>[1,1,1,3,3,4,3,2,4,2]
<strong>输出:</strong> true</pre>
</div>

####  解题思路
1. 哈希表记录出现过的元素 时间复杂O(n) 空间O(n)
2. 利用位操作技巧 时间复杂O(n) 空间O(1)

#### 题解
```go
// func containsDuplicate(nums []int) bool {
//     if len(nums) <= 1 {
//         return false
//     }
//     m := make(map[int]bool)
//     for _, i := range nums {
//         if m[i] {
//             return true
//         }
//         m[i] = true
//     }
//     return false
// }

func containsDuplicate(nums []int) bool {
	mark := 0
	for _, i := range nums {
		if mark & (1 << i) != 0 {
			return true
		}
		mark = mark | (1 << i)
	}
	return false
}
```