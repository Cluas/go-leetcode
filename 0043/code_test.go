package leetcode

import (
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{"2", "3", "6"}, {"123", "456", "56088"},
	}

	for _, tt := range tests {
		got := multiply(tt.num1, tt.num2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("multiply(%v, %v) return %v, want %v", tt.num1, tt.num2, got, tt.want)
		}
	}

}
