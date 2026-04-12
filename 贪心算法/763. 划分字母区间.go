package main

func partitionLabels(s string) (ans []int) {
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		end = max(end, last[c-'a'])
		if end == i {
			ans = append(ans, end-start+1)
			start = i + 1
		}
	}
	return ans
}
