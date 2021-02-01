//　数列aの範囲kの最小値をO(n)で求める
//  蟻本の4.4 P.300をgoに移植
func slideMin(a []int, k int) []int {
	n := len(a)
	b := make([]int, n-k+1)
	s, t := 0, 0

	deq := make([]int, n)
	for i := 0; i < n; i++ {
		for s < t && a[deq[t-1]] >= a[i] {
			t--
		}
		deq[t] = i
		t++
		if i-k+1 >= 0 {
			b[i-k+1] = a[deq[s]]
			if deq[s] == i-k+1 {
				s++
			}
		}
	}
	return b
}

//　数列aの範囲kの最大値をO(n)で求める
func slideMax(a []int, k int) []int {
	n := len(a)
	b := make([]int, n-k+1)
	s, t := 0, 0

	deq := make([]int, n)
	for i := 0; i < n; i++ {
		for s < t && a[deq[t-1]] <= a[i] {
			t--
		}
		deq[t] = i
		t++
		if i-k+1 >= 0 {
			b[i-k+1] = a[deq[s]]
			if deq[s] == i-k+1 {
				s++
			}
		}
	}
	return b
}
