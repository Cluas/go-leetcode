package leetcode

import (
	"reflect"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{
			s:    "egg",
			t:    "add",
			want: true,
		},
		{
			s:    "foo",
			t:    "bar",
			want: false,
		},
		{
			s:    "paper",
			t:    "title",
			want: true,
		},
	}
	for _, tt := range tests {
		got := isIsomorphic(tt.s, tt.t)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("wordPattern(%v, %v), return %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
