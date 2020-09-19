package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	km := make(map[int]bool, k)
	var i, j int
	for i < len(nums) {
		for j < len(nums) && j-i <= k {
			if km[nums[j]] {
				return true
			} else {
				km[nums[j]] = true
			}
			j++
		}
		delete(km, nums[i])
		i++
	}
	return false
}
