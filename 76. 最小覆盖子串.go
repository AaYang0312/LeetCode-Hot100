package main

func isCovered(cntS, cntT []int) bool {
	for i := 'A'; i <= 'Z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {

	m, n := len(s), len(t)
	if m < n {
		return ""
	}
	tcount := make([]int, 128)
	scount := make([]int, 128)

	for _, ch := range t {
		tcount[ch]++
	}

	minLen := m + 1
	result := ""

	l := 0

	for r := 0; r < m; r++ {
		// 扩展右边界
		scount[s[r]]++

		// 收缩左边界，直到窗口不再覆盖t
		for isCovered(scount, tcount) {
			// 更新最小覆盖子串
			if r-l+1 < minLen {
				minLen = r - l + 1
				result = s[l : r+1]
			}
			// 收缩左边界
			scount[s[l]]--
			l++
		}
	}

	return result
}
