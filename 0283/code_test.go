package leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums  []int
		moved []int
	}{
		{nums: []int{0, 1, 0, 3, 12}, moved: []int{1, 3, 12, 0, 0}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		moveZeroes(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.moved) {
			t.Fatalf("moveZeros(%v) return %v, want %v", nums, tt.nums, tt.moved)
		}
	}

}
