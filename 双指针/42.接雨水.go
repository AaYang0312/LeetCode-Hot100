package main

func trap(height []int) int {
	cache := height
	sum := 0
	n := len(cache)
	if n <= 2 {
		return 0
	}
	left, right := 0, n-1
	lmax, rmax := 0, 0
	for left < right {
		lmax = max(lmax, cache[left])
		rmax = max(rmax, cache[right])
		if lmax < rmax {
			sum += lmax - cache[left]
			left++
		} else {
			sum += rmax - cache[right]
			right--

		}

	}
	return sum
}
