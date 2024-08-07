package main

import (
	"bufio"
	"fmt"
	"math"
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

var H, W int
var s []string
var cnt [][]int

var x = []int{1, 0, -1, 0}
var y = []int{0, 1, 0, -1}

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	n            int
}

// newDsu :
func newDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size :
func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups : original implement
func (d DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}
	cnt = make([][]int, H)
	for i := 0; i < H; i++ {
		cnt[i] = make([]int, W)
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				cnt[i][j] = 1
				continue
			}
			for k := 0; k < 4; k++ {
				nx := j + x[k]
				ny := i + y[k]
				if nx < 0 || nx >= W || ny < 0 || ny >= H {
					continue
				}
				if s[ny][nx] == '#' {
					cnt[i][j] = 1
					break
				}
			}
		}
	}

	uf := newDsu(H * W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if cnt[i][j] == 1 {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := i+y[k], j+x[k]
				if ni < 0 || ni >= H || nj < 0 || nj >= W {
					continue
				}
				if cnt[ni][nj] == 0 {
					uf.Merge(i*W+j, ni*W+nj)
				}
			}
		}
	}

	ans := 0
	for _, e := range uf.Groups() {
		m := make(map[int]bool)
		for _, v := range e {
			i, j := v/W, v%W
			if cnt[i][j] == 1 {
				break
			}
			for k := 0; k < 4; k++ {
				ni, nj := i+y[k], j+x[k]
				if ni < 0 || ni >= H || nj < 0 || nj >= W {
					continue
				}
				if cnt[ni][nj] == 1 {
					m[ni*W+nj] = true
				}
			}
		}
		ans = max(ans, len(m)+len(e))
		// out(e, m)
	}
	out(ans)
}
