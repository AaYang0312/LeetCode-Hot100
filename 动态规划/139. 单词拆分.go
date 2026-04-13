package main

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	// 初始化
	maxLen := 0
	for _, word := range wordDict {
		words[word] = true
		maxLen = max(maxLen, len(word))
	}
	// 记忆
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 0; i <= len(s); i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if f[j] && words[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}
