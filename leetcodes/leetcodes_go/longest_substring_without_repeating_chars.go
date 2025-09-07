package leetcodesgo

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maximumSize, currentLeft := 1, 0
	charIndexer := make(map[rune]int)

	for idx, char := range s {
		if foundIdx, foundBool := charIndexer[char]; foundBool && currentLeft < foundIdx {
			currentLeft = foundIdx + 1
		}
		charIndexer[char] = idx
		maximumSize = max(maximumSize, idx-currentLeft+1)
	}
	return maximumSize
}

// func main() {
// 	evaluations := []string{"abcabcbb", "pwwkew", "pwpwkew", "au", "abba"}
// 	for _, ss := range evaluations {
// 		length := lengthOfLongestSubstring(ss)
// 		fmt.Println(ss, length)
// 	}
// }
