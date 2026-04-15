package main

func longestValidParentheses(s string) int {
	stc := []int{-1}
	ans := 0
	for i, c := range s {
		if c == '(' {
			stc = append(stc, i)
		} else if c == ')' && len(stc) > 1 { // 匹配成功
			stc = stc[:len(stc)-1]
			ans = max(ans, i-stc[len(stc)-1])
		} else { //栈底只有红线，如-1
			stc[0] = i
		}
	}
	return ans
}
