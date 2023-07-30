package main

import (
	"fmt"
)

func main() {
  //7,12 e,a ebcbbec_e_cabb_a_cecbbcbe
	fmt.Print(validPalindrome("ebcbbececabbacecbbcbe"))
  fmt.Println(" = true")
}

func validPalindrome(s string) bool {
	size := len(s);
	if size <= 2 {
		return true;
	}


	for i := 0; i < size/2; i++ {
		if s[i] != s[size-1-i] {
      return (validPalindromeWithoutDeletion(s[i:size-1-i]) || validPalindromeWithoutDeletion(s[i+1:size-i]))
		}
	}

	return true
}

func validPalindromeWithoutDeletion(s string) bool {
	size := len(s);
	if size <= 1 {
		return true;
	}

	for i := 0; i < size/2; i++ {
		if s[i] != s[size-1-i] {
			return false;
		}
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
