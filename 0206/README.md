### 206. 反转链表
<div class="notranslate"><p>反转一个单链表。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2-&gt;3-&gt;4-&gt;5-&gt;NULL
<strong>输出:</strong> 5-&gt;4-&gt;3-&gt;2-&gt;1-&gt;NULL</pre>

<p><strong>进阶:</strong><br>
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？</p>
</div>

#### 解题思路
记录当前上一个还有下一个元素的位置
遍历反转

#### 题解
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var pre, cur *ListNode
    cur = head
    for cur != nil {
        pre, cur, cur.Next = cur, cur.Next, pre
    }
    return pre
}
```