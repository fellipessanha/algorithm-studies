package leetcodesgo

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingChars(t *testing.T) {
	evaluations := map[string]int{"abcabcbb": 4, "pwwkew": 4, "pwpwkew": 4, "au": 2, "abba": 2}
	for entry, expectedResult := range evaluations {
		result := lengthOfLongestSubstring(entry)
		if result != expectedResult {
			t.Errorf("Result is wrong! Evaluating %s: expected %d, got %d", entry, expectedResult, result)
		}
	}

}
