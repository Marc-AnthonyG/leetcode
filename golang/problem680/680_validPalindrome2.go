package main

import (
	"fmt"
)

func main() {
	//7,12 e,a ebcbbec_e_cabb_a_cecbbcbe
	fmt.Print(validPalindrome("ebcbbececabbacecbbcbe"))
	fmt.Println(" = true")

	fmt.Print(validPalindrome("abc"))
	fmt.Println(" = false")
}

func validPalindrome(s string) bool {
  var size int;
  size = len(s)/2

	for i := 0; i < size/2; i++ {
		if s[i] != s[size-1-i] {
      return (validPalindromeWithoutDeletion(s, i, size-2-i) || validPalindromeWithoutDeletion(s, i+1, size-1-i ))
		}
	}

	return true
}

func validPalindromeWithoutDeletion(s string, left int, right int) bool {
	for left < right {
    if s[left] != s[right]{
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
