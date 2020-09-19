package leetcode

import (
	"reflect"
	"testing"
)

func TestContainsNearbyDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
	}
	for _, tt := range tests {
		got := containsNearbyDuplicate(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("containsNearbyDuplicate(%v, %v), returns %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}

}
