package leetcode

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
