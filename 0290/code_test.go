package leetcode

import (
	"reflect"
	"testing"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern string
		s       string
		want    bool
	}{
		{
			pattern: "abba",
			s:       "dog cat cat dog",
			want:    true,
		},
		{
			pattern: "abba",
			s:       "dog cat cat fish",
			want:    false,
		},
		{
			pattern: "aaaa",
			s:       "dog cat cat dog",
			want:    false,
		},
		{
			pattern: "abba",
			s:       "dog dog dog dog",
			want:    false,
		},
	}
	for _, tt := range tests {
		got := wordPattern(tt.pattern, tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("wordPattern(%v, %v), return %v, want %v", tt.pattern, tt.s, got, tt.want)
		}
	}
}
