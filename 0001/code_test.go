package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 3, 4, 6}, 6, []int{0, 2}},
		{[]int{2, 3, 4, 7}, 8, nil},
		{nil, 8, nil},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("twoSum(%v, %v) return %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
