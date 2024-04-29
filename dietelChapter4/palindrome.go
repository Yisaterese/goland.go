package main

import (
	"fmt"
	"strings"
)

func palindrome(newPalindrome string) bool {

	newPalindrome = strings.ToLower(newPalindrome)
	for i := 0; i < len(newPalindrome); i++ {
		for j := len(newPalindrome) - 1; j > 0; j-- {
			if newPalindrome[i] == newPalindrome[j] {
				return true
			}
		}
	}

	return false

}
func main() {
	fmt.Println(palindrome("park"))
}
