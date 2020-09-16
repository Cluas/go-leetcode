package leetcode

import (
	"reflect"
	"testing"
)

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		A    []int
		B    []int
		C    []int
		D    []int
		want int
	}{
		{
			A:    []int{1, 2},
			B:    []int{-2, -1},
			C:    []int{-1, 2},
			D:    []int{0, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := fourSumCount(tt.A, tt.B, tt.C, tt.D)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("wordPattern(%v, %v, %v, %v), return %v, want %v", tt.A, tt.B, tt.C, tt.D, got, tt.want)
		}
	}
}
