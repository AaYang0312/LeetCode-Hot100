package main

func topKFrequent(nums []int, k int) (ans []int) {
	cnt := make(map[int]int, len(nums))
	maxCnt := 0
	for _, n := range nums {
		cnt[n]++
		maxCnt = max(maxCnt, cnt[n])
	}
	buckets := make([][]int, maxCnt+1)
	for i, c := range cnt {
		buckets[c] = append(buckets[c], i)
	}
	for i := maxCnt; len(ans) < k; i-- {
		ans = append(ans, buckets[i]...)
	}
	return
}
