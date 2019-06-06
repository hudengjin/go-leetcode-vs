/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
    for index, value := range nums {
		if value == target {
			return index
		} else if  index < len(nums) - 1 {
			if target > value && target < nums[index + 1]{
				return index + 1
			}
		} else {
			return index + 1
		}
	}
	return 0
	// if len(nums) == 0 {
	// 	return 0
	// }
	// left := 0
	// right := len(nums) - 1
	// for left <= right {
	// 	mid := (left + right) / 2
	// 	if nums[mid] == target {
	// 		return mid
	// 	} else if target < nums[mid] {
	// 		right = mid - 1
	// 	} else {
	// 		left = mid + 1
	// 	}
	// }
	// return left
	// return sort.SearchInts(nums, target)
}

