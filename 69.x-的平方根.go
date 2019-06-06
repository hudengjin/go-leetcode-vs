/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
func mySqrt(x int) int {
	// return int(math.Sqrt(float64(x)))
	//return int(NewtonSqrt(float64(x)))
	return BinarySqrt(x)
}

func NewtonSqrt(num float64) float64 {
	x := num / 2.0
	var y float64 = 0
	count := 1
	for math.Abs(x-y) > 1e-7 {
		count += 1
		y = x
		x = (1.0/2.0)*x + (num*1.0)/(x*2.0)
	}
	return x
}

func BinarySqrt(num int) int {
	low, high := 0, num
	mid := num / 2 + 1
	for low <= high {
		mid = low + (high - low ) >> 1
		if mid*mid > num {
			high = mid - 1
		} else {
			if mid == num || (mid+1)*(mid+1) > num {
				return mid
			}
			low = mid + 1
		}
	}
	return mid
}

