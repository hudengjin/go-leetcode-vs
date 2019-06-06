/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
func lengthOfLastWord(s string) int {
    if len(s) == 0 {
		return 0
	}
	strArray := strings.Split(s, " ")
	if strArray[len(strArray) - 1] != "" {
		return len(strArray[len(strArray) - 1])
	} else {
		for j := len(strArray) - 2; j >= 0; j-- {
			if strArray[j] != ""{
				return len(strArray[j])
			}
		}
	}
	return 0
}

