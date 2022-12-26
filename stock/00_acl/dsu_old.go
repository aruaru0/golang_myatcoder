type UnionFind struct {
	n   int
	par []int
}

func newUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.n = n
	uf.par = make([]int, n)
	for i, _ := range uf.par {
		uf.par[i] = -1
	}
	return uf
}
func (uf UnionFind) root(x int) int {
	if uf.par[x] < 0 {
		return x
	}
	uf.par[x] = uf.root(uf.par[x])
	return uf.par[x]
}
func (uf UnionFind) unite(x, y int) {
	rx, ry := uf.root(x), uf.root(y)
	if rx != ry {
		if uf.size(rx) > uf.size(ry) {
			rx, ry = ry, rx
		}
		uf.par[ry] += uf.par[rx]
		uf.par[rx] = ry
	}
}
func (uf UnionFind) same(x, y int) bool {
	return uf.root(x) == uf.root(y)
}
func (uf UnionFind) size(x int) int {
	return -uf.par[uf.root(x)]
}
func (uf UnionFind) groups() [][]int {
	rootBuf, groupSize := make([]int, uf.n), make([]int, uf.n)
	for i := 0; i < uf.n; i++ {
		rootBuf[i] = uf.root(i)
		groupSize[rootBuf[i]]++
	}
	res := make([][]int, uf.n)
	for i := 0; i < uf.n; i++ {
		res[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < uf.n; i++ {
		res[rootBuf[i]] = append(res[rootBuf[i]], i)
	}
	result := make([][]int, 0, uf.n)
	for i := 0; i < uf.n; i++ {
		if len(res[i]) != 0 {
			r := make([]int, len(res[i]))
			copy(r, res[i])
			result = append(res, r)
		}
	}
	return result
}
