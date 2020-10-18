package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

var h, w, sx, sy, gx, gy int
var b [][]int
var d [][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	h, w = getI(), getI()
	sy, sx, gy, gx = getI()-1, getI()-1, getI()-1, getI()-1
	b = make([][]int, h)
	for i := 0; i < h; i++ {
		b[i] = make([]int, w)
		s := getS()
		for j := 0; j < w; j++ {
			b[i][j] = int(s[j] - '0')
		}
	}
	uf := newUnionFind(h * w)
	for y := 0; y < h; y++ {
		for x := 0; x < w-1; x++ {
			if abs(b[y][x]-b[y][x+1]) <= 1 {
				uf.unite(y*w+x, y*w+(x+1))
			}
		}
		for x := 0; x < w-2; x++ {
			if b[y][x] == b[y][x+2] && b[y][x] > b[y][x+1] {
				uf.unite(y*w+x, y*w+(x+2))
			}
		}
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h-1; y++ {
			if abs(b[y][x]-b[y+1][x]) <= 1 {
				uf.unite(y*w+x, (y+1)*w+x)
			}
		}
		for y := 0; y < h-2; y++ {
			if b[y][x] == b[y+2][x] && b[y][x] > b[y+1][x] {
				uf.unite(y*w+x, (y+2)*w+x)
			}
		}
	}

	if uf.same(sy*w+sx, gy*w+gx) {
		out("YES")
		return
	}
	out("NO")
}

// UnionFind高速
type UnionFind struct {
	d []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.d = make([]int, N)
	for i := 0; i < N; i++ {
		u.d[i] = -1
	}
	return u
}

func (p *UnionFind) root(x int) int {
	if p.d[x] < 0 {
		return x
	}
	p.d[x] = p.root(p.d[x])
	return p.d[x]
}

func (p *UnionFind) unite(x, y int) bool {
	x = p.root(x)
	y = p.root(y)
	if x == y {
		return false
	}
	if p.d[x] > p.d[y] {
		x, y = y, x
	}
	p.d[x] += p.d[y]
	p.d[y] = x
	return true
}

func (p *UnionFind) same(x, y int) bool {
	return p.root(x) == p.root(y)
}

func (p *UnionFind) size(x int) int {
	return -p.d[p.root(x)]
}
