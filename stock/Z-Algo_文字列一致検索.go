// 先頭からの文字列と各文字から一致する文字列の長さを調べる
// 戻り値は各文字からの一致文字数（０は文字列の長さと一致）
func zalgo(str string) []int {
	n := len(str)
	a := make([]int, n)
	from, last := -1, -1
	a[0] = n
	for i := 1; i < n; i++ {
		idx := a[i]
		if from != -1 {
			idx = min(a[i-from], last-i)
			idx = max(0, idx)
		}
		for idx+i < n && str[idx] == str[idx+i] {
			idx++
		}
		a[i] = idx
		if last < i+idx {
			last = i + idx
			from = i
		}
	}
	return a
}