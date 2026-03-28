package main

func groupAnagrams(strs []string) [][]string {
	// 创建一个map存储26个字母出现次数
	m := make(map[[26]int][]string)
	// 统计strs切片中出现的所有词语
	for _, verbal := range strs {
		cnt := [26]int{}
		// 遍历每一个字符
		for _, c := range verbal {
			// 统计加一
			cnt[c-'a']++
		}
		// 计入map
		m[cnt] = append(m[cnt], verbal)
	}
	// 结果，最长长度可能与strs相同
	ans := make([][]string, 0, len(strs))
	// 将所有分组的字符串切片添加到ans
	for _, v := range m {
		ans = append(ans, v)
	}
	// 返回分组结果
	return ans
}
