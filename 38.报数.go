/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 */
func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = nextTime(res)
	}
	return res
}

func nextTime(res string) string {
	i := 0
	// str := []rune(res)
	n := len(res)
	temp := ""
	for i < n {
		count := 1
		for i < n - 1 && res[i] == res[i + 1] {
				i++
				count++
		}
		temp += strconv.Itoa(count) + string(res[i])
		i++
	}
	return temp

}

