package leetcode

import (
	"reflect"
	"testing"
)

func TestIshappy(t *testing.T) {

	tests := []struct {
		n    int
		want bool
	}{
		{19, true}, {2, false},
	}

	for _, tt := range tests {
		got := ishappy(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("ishappy(%v) return %v want %v", tt.n, got, tt.want)
		}
	}
}
