/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
func maxSubArray(nums []int) int {
	sum := 0
	res := nums[0]
	for _, num := range nums {
		// sum = sum > 0 ? sum + num : num
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if res < sum {
			res = sum
		} 
		// res = math.Max(res, sum)
	}
	return res
}

