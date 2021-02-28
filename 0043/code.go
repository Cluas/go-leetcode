package leetcode

import "strings"

func multiply(num1 string, num2 string) string {
	var ints = make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			sum := int(mul) + ints[i+j+1]
			ints[i+j+1] = sum % 10
			ints[i+j] += sum / 10
		}
	}

	var idx int
	for idx < len(ints)-1 && ints[idx] == 0 {
		idx++
	}
	ints = ints[idx:]

	var buf strings.Builder
	buf.Grow(len(ints))
	for _, i := range ints {
		buf.WriteByte('0' + byte(i))
	}

	return buf.String()
}
