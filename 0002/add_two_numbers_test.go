package add_two_numbers

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func New(vv ...int) *ListNode {
	l := new(ListNode)
	n := l
	for _, v := range vv {
		n.Next = &ListNode{Val: v}
		n = n.Next
	}
	return l.Next
}

func (l *ListNode) String() string {
	if l == nil {
		return "<nil>"
	}

	var b strings.Builder

	b.WriteByte('(')
	b.WriteString(strconv.Itoa(l.Val))
	for node := l.Next; node != nil; node = node.Next {
		b.WriteString(" -> ")
		b.WriteString(strconv.Itoa(node.Val))
	}
	b.WriteByte(')')

	return b.String()
}
func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1, l2 *ListNode
		sum    *ListNode
	}{
		{New(2, 4, 6), New(5, 6, 4), New(7, 0, 1, 1)},
		{New(2, 4, 3), New(5, 6, 4), New(7, 0, 8)},
		{New(2, 4), New(5, 6, 4), New(7, 0, 5)},
		{New(2), New(5, 6, 4), New(7, 6, 4)},
		{New(2), New(5, 6), New(7, 6)},
		{New(2), New(5), New(7)},
	}

	for _, tt := range tests {
		sum := addTwoNumbers(tt.l1, tt.l2)
		if reflect.DeepEqual(sum, tt.sum) == false {
			t.Errorf("addTwoNumbers(%v, %v) return %v, want %v", tt.l1, tt.l2, sum, tt.sum)
		}
	}
}
