package leetcode

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}
	var l, r int
	for l < len(nums) {
		r = l + 1
		for r < len(nums) && r-l <= k {
			if abs(nums[r]-nums[l]) <= t {
				return true
			}
			r++
		}
		l++
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
