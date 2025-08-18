package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	fmt.Println("str:", str)
	reversed := ""
	for _, charvalue := range str {
		//fmt.Println("char:",string(charvalue))
		reversed = string(charvalue) + reversed
		//fmt.Println("reversed:",reversed)
	}

	return str == reversed
}

func main() {
	fmt.Println("121 is palindrome:", isPalindrome(121))     // true
	fmt.Println("-121 is palindrome:", isPalindrome(-121))   // true
	fmt.Println("10 is palindrome:", isPalindrome(10))       // true
	fmt.Println("0 is palindrome:", isPalindrome(0))         // true
	fmt.Println("12321 is palindrome:", isPalindrome(12321)) // true
}
