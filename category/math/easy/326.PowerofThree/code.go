package leetcode

func isPowerOfThree(n int) bool {
	if (n%3 != 0 && n != 1) || n == 0 {
		return false
	}
	for n != 0 {
		n /= 3
		if n%3 != 0 && n != 1 {
			return false
		}
	}
	return true
}
