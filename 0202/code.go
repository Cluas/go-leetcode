package leetcode

func ishappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1

}

func step(n int) int {
	var squares = [10]int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}

	var res int
	for ; n > 0; n /= 10 {
		res += squares[n%10]
	}
	return res
}
