package ReverseInteger

import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		mod := x % 10
		if res > math.MaxInt32/10 || res < math.MinInt32/10{
			return 0
		}
		if (res == math.MaxInt32/10 && mod > 7) || (res == math.MinInt32/10 && mod < -8) {
			return 0
		}
		res = res * 10 + mod
		x = x / 10
	}
	return res
}