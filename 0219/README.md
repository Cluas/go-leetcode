### 219. 存在重复元素 II
<div class="notranslate"><p>给定一个整数数组和一个整数&nbsp;<em>k</em>，判断数组中是否存在两个不同的索引<em>&nbsp;i</em>&nbsp;和<em>&nbsp;j</em>，使得&nbsp;<strong>nums [i] = nums [j]</strong>，并且 <em>i</em> 和 <em>j</em>&nbsp;的差的 <strong>绝对值</strong> 至多为 <em>k</em>。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> nums = [1,2,3,1], k<em> </em>= 3
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>nums = [1,0,1,1], k<em> </em>=<em> </em>1
<strong>输出:</strong> true</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入: </strong>nums = [1,2,3,1,2,3], k<em> </em>=<em> </em>2
<strong>输出:</strong> false</pre>
</div>

### 解题思路
![image.png](https://pic.leetcode-cn.com/1600514512-qPGOLw-image.png)
先构造一个长度为k的字典，用来判定是否有重复元素
开始遍历i 然后逐步加大j 若找到重复 返回ture 否则滑动窗口继续找 当j到底后还没找到 返回false


### 代码

```golang
func containsNearbyDuplicate(nums []int, k int) bool {
    km := make(map[int]bool, k)
    var i, j int
    for j < len(nums) {
        for j < len(nums) && j - i  <= k  {
            if km[nums[j]] {
                return true
            } else {
                km[nums[j]] = true
            }
            j++
        }
        delete(km, nums[i])
        i++
    }
    return false
}
```