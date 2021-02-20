package NumberOf1Bits

func hammingWeight(num uint32) int {
	var c int
	for num > 0 {
		c++
		num &= num - 1
	}
	return c
}