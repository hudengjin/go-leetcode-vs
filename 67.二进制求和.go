/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
func addBinary(a string, b string) string {
	// iA := binaryStr2Int(a)
	// iB := binaryStr2Int(b)
	// return int2BinStr(iA + iB)
	var carry, sum int
	i, j := len(a), len(b)
	if i < j {
		i, j = j, i
		a, b = b, a
	}
	res := make([]byte, i + 1)
	for j > 0 {
		j--
		i--
		sum = int(a[i] - '0') + int(b[j] - '0') + carry
		carry = sum / 2
		sum = sum % 2
		res[i+1] = byte(sum + '0')
	}
	for i > 0 {
		i--
		sum = int(a[i] - '0') + carry
		carry = sum / 2
		sum = sum % 2
		res[i+1] = byte(sum + '0')
	}
	res[0] = byte(carry + '0')
	for i < len(res) - 1 {
		if res[i] == '0' {
			i++
		} else {
			break
		}
	}
	return string(res[i:])
}

// func binaryStr2Int(s string) int {
// 	l := len(s)
// 	num := 0
// 	for i := l - 1; i >= 0; i-- {
// 		num += (int(s[l-i-1]) & 0xf) << uint(i)
// 	}
// 	return num
// }

// func int2BinStr(n int) string {
// 	if n == 0 {
// 		return "0"
// 	}
// 	s := ""
// 	for q := n; q > 0; q /= 2 {
// 		m := q % 2
// 		s = fmt.Sprintf("%v%v", m, s)
// 	} 
// 	return s
// }
