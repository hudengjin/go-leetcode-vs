/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
func climbStairs(n int) int {
    return  climbStairs1(n)
}

func climbStairs1(n int ) int {
	if n <= 2 {
		return n
	}
	first := 2
	second := 3
	for i := 3; i < n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func climbStairs2(n int)  int {
	if n <= 2 {
		return n
	}
	count := []int{1, 2}
	for i := 3; i <= n; i++ {
		count = append(count, count[i-2] + count[i-3])
	}
	return count[len(count) - 1]
}

func climbStairs3(i, n int) int {
	if i > n {
		return 0
	}
	if (i == n) {
		return 1
	}
	return climbStairs3(i+1, n) + climbStairs3(i+2, n)
}
