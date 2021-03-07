package leetcode

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func moveZeroes2(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if i > j {
			nums[i] = 0
		}
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
}
