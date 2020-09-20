package leetcode

// func containsDuplicate(nums []int) bool {
//     if len(nums) <= 1 {
//         return false
//     }
//     m := make(map[int]bool)
//     for _, i := range nums {
//         if m[i] {
//             return true
//         }
//         m[i] = true
//     }
//     return false
// }

func containsDuplicate(nums []int) bool {
	mark := 0
	for _, i := range nums {
		if mark&(1<<i) != 0 {
			return true
		}
		mark = mark | (1 << i)
	}
	return false
}
