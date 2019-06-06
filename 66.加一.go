/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
func plusOne(digits []int) []int {
	
	for j := len(digits) - 1; j >= 0; j-- {
		if digits[j] < 9 {
			digits[j]++
			return digits
		} else {
			digits[j] = 0
		}
	}
	return append([]int{1}, digits...) 
}

