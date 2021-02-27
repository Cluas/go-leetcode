package leetcode

import (
	"reflect"
	"testing"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		{
			"horse",
			"ros",
			3,
		},
		{
			"intention",
			"execution",
			5,
		},
	}

	for _, tt := range tests {
		got := minDistance(tt.word1, tt.word2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("minDistance(%v, %v) return %v, want %v", tt.word1, tt.word2, got, tt.want)
		}
	}
}
