package leetcode

import (
	"reflect"
	"testing"
)

func TestTrailingZeros(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 0,
		},
		{
			n:    5,
			want: 1,
		},
	}
	for _, tt := range tests {
		got := trailingZeroes(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("trailingZeroes(%v), returns %v, want %v", tt.n, got, tt.want)
		}
	}
}
