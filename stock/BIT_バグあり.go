type bit []int

func (p *bit) init(n int) {
	*p = make([]int, n+1)
}
func (p bit) sum(i int) int {
	if i == 0 {
		return 0
	}
	return p[i] + p.sum(i-(i&-i))
}
func (p bit) add(i, x int) {
	if i >= len(p) {
		return
	}
	p[i] += x
	p.add(i+(i&-i), x)
}