package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		// если символ уже был и он внутри окна
		if idx, ok := lastIndex[ch]; ok && idx >= left {
			left = idx + 1
		}

		lastIndex[ch] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen

}
