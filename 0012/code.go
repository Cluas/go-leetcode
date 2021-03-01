package leetcode

import "strings"

func intToRoman(num int) string {
	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var b strings.Builder
	for i, n := range ints {
		for num >= n {
			num -= n
			b.WriteString(romans[i])
		}
	}
	return b.String()
}
