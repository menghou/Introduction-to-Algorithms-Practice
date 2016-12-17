package ch02

func BinarySum(a, b int) int {
	if b == 0 {
		return a
	} else {
		return BinarySum((a ^ b), (a&b)<<1)
	}
}
