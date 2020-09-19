package leetcode

import (
	"reflect"
	"testing"
)

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			want:   3,
		},
		{
			points: [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			want:   4,
		},
	}
	for _, tt := range tests {
		got := maxPoints(tt.points)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxPoints(%v), returns %v, want %v", tt.points, got, tt.want)
		}
	}

}
