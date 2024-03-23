package longest_substring

func lengthOfLongestSubstring(s string) int {
	max := 0
	chars := make(map[rune]int, len(s))
	startString := 0

	for k, v := range s {
		if lastIndex, isValid := chars[v]; isValid && lastIndex >= startString {
			startString = lastIndex + 1
		}
		chars[v] = k

		if k-startString+1 > max {
			max = k - startString + 1
		}

	}
	return max
}
