package leetcode

import (
	"reflect"
	"testing"
)

func TestSumSubarrayMins(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{3, 1, 2, 4}, 17}, {[]int{11, 81, 94, 43, 3}, 444},
	}

	for _, tt := range tests {
		got := sumSubarrayMins(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("sumSubarrayMins(%v) return %v, want %v", tt.arr, got, tt.want)
		}
	}

}
