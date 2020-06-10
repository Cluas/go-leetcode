package longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("lengthOfLongestSubstring(%v) got %v, want %v", tt.s, got, tt.want)
		}
	}
}
