package main

import (
	"strings"
	"unicode"
)

func main() {
	println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			if left >= right {
				return true
			}
		}
		for !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			if left >= right {
				return true
			}
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

/**
  Task: decide if a string is a palindrome after deleting zero or one character

  Note for this problem:
  So as i see it, the letter at the center for word with odd number of letter
  doesn't matter. So you can just check if the first letter equal the last etc.

  It does get triky with deleting a letter
  Test -> tet -> delete close to the center

  ABBCA -> ABBA same case ACBBA -> ABBA

  ABCCBBA-> ABBA same case ACBBA -> ABBA

  ebcbbececabbacecbbcbe  -> ebcbb_e_cecabbace_c_bbcbe

  si on delete c ca marche pas
*/
