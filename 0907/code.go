package leetcode

func sumSubarrayMins(arr []int) int {
	stack, dp, res, mod := []int{}, make([]int, len(arr)+1), 0, 1000000007
	stack = append(stack, -1)
	for i := 0; i < len(arr); i++ {
		for stack[len(stack)-1] != -1 && arr[i] <= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*arr[i]
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
	}
	return res
}
