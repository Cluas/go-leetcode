package leetcode

import (
	"reflect"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "tree",
			want: "eetr",
		},
		{
			s:    "cccaaa",
			want: "cccaaa",
		},
		{
			s:    "Aabb",
			want: "bbAa",
		},
	}
	for _, tt := range tests {
		got := frequencySort(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("wordPattern(%v,), return %v, want %v", tt.s, got, tt.want)
		}
	}
}
