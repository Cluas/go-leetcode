package leetcode

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
	}
	for _, tt := range tests {
		got := findErrorNums(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findErrorNums(%v), return %v, want %v", tt.nums, got, tt.want)
		}
	}
}
