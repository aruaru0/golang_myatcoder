func nCk(n, k int) int {
	if k < 0 {
		return 0
	}
	if k == 0 {
		return 1
	}
	tot := 1
	for i := 1; i <= k; i++ {
		tot = tot * (n + 1 - i) / i
	}
	return tot
}