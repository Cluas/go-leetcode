package leetcode

import (
	"reflect"
	"testing"
)

func TestNumberOfBoomerangs(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{points: [][]int{{0, 0}, {1, 0}, {2, 0}}, want: 2},
	}

	for _, tt := range tests {
		got := numberOfBoomerangs(tt.points)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("numberOfBoomerangs(%v), returns %v, want %v", tt.points, got, tt.want)
		}
	}
}
