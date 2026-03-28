package main

func findAnagrams(s string, p string) []int {
	ls, lp := len(s), len(p)
	result := []int{}
	if ls < lp {
		return []int{}
	}
	var scount, pcount [26]int
	for i, ch := range p {
		pcount[ch-'a']++
		scount[s[i]-'a']++
	}
	if scount == pcount {
		result = append(result, 0)
	}
	left := 0

	for right := lp; right < ls; right++ {

		left = right - lp

		scount[s[left]-'a']--
		scount[s[right]-'a']++
		if scount == pcount {
			result = append(result, left+1)
		}
	}
	return result
}
