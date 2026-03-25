package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charMap := map[byte]int{}
	left, max := 0, 1
	charMap[s[0]] = 0
	for right := 1; right < len(s); right++ {
		if idx, ok := charMap[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		charMap[s[right]] = right
		if right-left+1 > max {
			max = right - left + 1
		}
	}
	return max
}
