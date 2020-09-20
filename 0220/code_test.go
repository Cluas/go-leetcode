package leetcode

import (
	"reflect"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		t    int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			t:    0,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			t:    2,
			want: true,
		},
		{
			nums: []int{1, 5, 9, 1, 5, 9},
			k:    2,
			t:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		got := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("containsNearbyDuplicate(%v, %v), returns %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}

}
