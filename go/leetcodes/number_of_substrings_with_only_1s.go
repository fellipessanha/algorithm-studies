package leetcodesgo

import "strings"

func numSub(s string) int {
	largestOneSubstrings := strings.Split(s, "0")
	counter := 0
	for _, substring := range largestOneSubstrings {
		substringLength := len(substring)
		counter = (counter + substringLength*(substringLength+1)/2) % (1e9 + 7)
	}
	return counter
}
