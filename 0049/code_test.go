package leetcode

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	var tests = []struct {
		ss   []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}

	for _, tt := range tests {

		got := groupAnagrams(tt.ss)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("groupAnagrams(%v) return %v, want %v", tt.ss, got, tt.want)
		}
	}
}
