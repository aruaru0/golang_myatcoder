
// ｎ進数を計算（-2進数などに対応）
func calcN(x, n int) []int {
	ret := make([]int, 0)
	for x != 0 {
		r := x % n
		if r < 0 {
			r += (-n)
		}
		x = (x - r) / n
		ret = append(ret, r)
	}
	if len(ret) == 0 {
		ret = append(ret, 0)
	}
	return ret
}