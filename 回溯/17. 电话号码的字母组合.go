package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var mapping = []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	path := make([]byte, len(digits))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			res = append(res, string(path)) // string 进行内存拷贝，不需要额外设定tmp
			return
		}
		for _, c := range mapping[digits[i]-'0'] {
			path[i] = byte(c)
			dfs(i + 1)
		}
	}
	dfs(0)
	return res
}
