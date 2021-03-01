package leetcode

import (
	"reflect"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{3, "III"}, {4, "IV"}, {9, "IX"}, {58, "LVIII"}, {1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		got := intToRoman(tt.num)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("intToRoman(%v) return %v, got %v", tt.num, got, tt.want)
		}
	}

}
