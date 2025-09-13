package leetcodesgo

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	result := longestPalindrome("oi eu sou o abba goku")
	expectedResult := " abba "

	if result != expectedResult {
		t.Errorf("Wrong answer! expected '%s', got '%s'\n", expectedResult, result)
	}

}
