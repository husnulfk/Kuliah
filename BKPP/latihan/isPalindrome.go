package main

import "fmt"

func main() {

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1221))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0

	for x > revertedNumber {
		digit := x % 10
		revertedNumber = (revertedNumber * 10) + digit
		x = x / 10
	}
	return x == revertedNumber || x == revertedNumber/10
}
