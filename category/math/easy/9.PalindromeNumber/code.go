package math

import "strconv"

func isPalindrome(x int) bool {

	xString := strconv.Itoa(x)
	var isPalindromeV = true

	for i := 0; i < len(xString); i++ {
		if (len(xString)-i-1)-i <= 0 {
			break
		}
		if xString[i] != xString[len(xString)-i-1] {
			isPalindromeV = false
			return isPalindromeV
		}
	}
	return isPalindromeV
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	result, newX := 0, x
	for newX > 0 {
		number := newX % 10
		result = result*10 + number
		newX /= 10
	}
	return result == x
}

