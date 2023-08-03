package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	return (tmp == ans)
}

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome1(x int) bool {
	// convert int to slice of runes
	runes := []rune(strconv.Itoa(x))

	// reverse the runes in the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// convert back to int, via string
	y, err := strconv.Atoi(string(runes))

	// check for an error (overflow) and equality
	if err == nil && x == y {
		return true
	}
	return false
}
